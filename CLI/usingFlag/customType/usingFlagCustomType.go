package main

import (
	"flag"
	"fmt"
	"net/url"
	"reflect"
)

type Value interface {
	String() string
	Set(string) error
}

type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

var u = &url.URL{}

func init() {
	flag.Var(&URLValue{u}, "url", "URL to parse")
}

func main() {
	flag.Parse()
	if reflect.ValueOf(*u).IsZero() {
		panic("did not pass an URL")
	}
	fmt.Printf(`{scheme: %q, host: %q, path: %q`, u.Scheme, u.Host, u.Path)
}
