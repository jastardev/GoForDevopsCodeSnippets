package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

type getReq struct {
	Author string `json:"author"`
}

func (g *getReq) fromReader(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, g)
}

type getResp struct {
	Quote string `json:"quote"`
	Error *Error `json:"error"`
}

type ErrCode string

const (
	// UnknownCode indicates the ErrCode was not set, aka the zero value.
	UnknownCode ErrCode = ""
	// UnknownAuthor indicates that the request wanted a quote from an
	// author that didn't exist in the server.
	UnknownAuthor ErrCode = "UnknownAuthor"
)

type Error struct {
	Code ErrCode
	Msg  string
}

func (e *Error) Error() string {
	return fmt.Sprintf("(code %v): %s", e.Code, e.Msg)
}

type server struct {
	serv   *http.Server
	quotes map[string][]string
}

func newServer(port int) (*server, error) {
	s := &server{
		serv: &http.Server{
			Addr: ":" + strconv.Itoa(port),
		},
		quotes: map[string][]string{
			"Mark Twain": {
				"History doesn't repeat itself, but it does rhyme",
				"Lies, damned lies, and statistics",
				"Golf is a good walk spoiled",
			},
			"Benjamin Franklin": {
				"Tell me and I forget. Teach me and I remember. Involve me and I learn",
				"I didn't fail the test. I just found 100 ways to do it wrong",
			},
			"Eleanor Roosevelt": {
				"The future belongs to those who believe in the beauty of their dreams",
			},
		},
	}
	mux := http.NewServeMux()
	mux.HandleFunc(`/qotd/v1/get`, s.qotdGet)
	s.serv.Handler = mux
	return s, nil
}
func (s *server) start() error {
	return s.serv.ListenAndServe()
}

func (s *server) qotdGet(w http.ResponseWriter, r *http.Request) {
	req := getReq{}
	if err := req.fromReader(r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var quotes []string

	if req.Author == "" {
		for _, quotes = range s.quotes {
			break
		}
	} else {
		var ok bool
		quotes, ok = s.quotes[req.Author]
		if !ok {
			b, err := json.Marshal(
				getResp{
					Error: &Error{
						Code: UnknownAuthor,
						Msg:  fmt.Sprintf("Author %q was not found", req.Author),
					},
				},
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.Write(b)
			return
		}
	}

	i := rand.Intn(len(quotes))
	b, err := json.Marshal(getResp{Quote: quotes[i]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(b)
	return
}

func main() {
	serv, err := newServer(80)
	if err != nil {
		panic(err)
	}
	// Start our server. This blocks, so we have it do it in its own goroutine.
	serv.start()

}
