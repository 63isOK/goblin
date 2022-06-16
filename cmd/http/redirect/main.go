package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
	c := http.DefaultClient
	isRedirect := false
	var rurl *url.URL
	c.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		isRedirect = true
		rurl = req.URL
		return nil
	}
	c.Timeout = 60 * time.Second
	r, err := http.NewRequest("GET", "http://8.134.75.117/v2/server_api/query/record/image?path=imgs/b4bd6dfe-e26a6f9c/20220616/1655363800_112391470_b4bd6dfe-e26a6f9c_0_36797212_T247_full.jpg", nil)
	if err != nil {
		panic(err)
	}
	rs, err := c.Do(r)
	// r, err := http.Get("http://baidu.com")
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	body, err := io.ReadAll(rs.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(len(body))

	if isRedirect {
		r, err := http.NewRequest("GET", rurl.String(), nil)
		if err != nil {
			panic(err)
		}
		rs, err := c.Do(r)
		// r, err := http.Get("http://baidu.com")
		if err != nil {
			panic(err)
		}
		defer rs.Body.Close()

		body, err := io.ReadAll(rs.Body)
		if err != nil {
			panic(err)
		}
		fmt.Print(len(body))
	}
}
