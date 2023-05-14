package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type getReq struct {
	Author string `json:"author"`
}

type getResp struct {
	Quote string `json:"quote"`
	Error *Error `json:"error"`
}

type ErrorCode string

type Error struct {
	Code ErrorCode
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("(code %v): %s", e.Code, e.Msg)
}

type QOTD struct {
	u      *url.URL
	client *http.Client
}

func New(addr string) (*QOTD, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	return &QOTD{u: u, client: &http.Client{}}, nil
}

func (q *QOTD) restCall(ctx context.Context, endpoint string, req, resp interface{}) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 2*time.Second)
		defer cancel()
	}
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	hReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	hResp, err := q.client.Do(hReq)
	if err != nil {
		return err
	}
	b, err = io.ReadAll(hResp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, resp)
}

func (q *QOTD) Get(ctx context.Context, author string) (string, error) {
	const endpoint = `/qotd/v1/get`
	ref, _ := url.Parse(endpoint)

	resp := getResp{}

	err := q.restCall(ctx, q.u.ResolveReference(ref).String(), getReq{Author: author}, &resp)
	switch {
	case err != nil:
		return "", err
	case resp.Error != nil:
		return "", resp.Error
	}
	return resp.Quote, nil
}

func main() {
	client, err := New("http://localhost:80")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	quote, err := client.Get(ctx, "Mark Twain")
	if err != nil {
		panic(err)
	}

	fmt.Println(quote)
}
