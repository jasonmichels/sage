package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const (
	SERVER1 = "https://www.fisdap.net/"
	SERVER2 = "http://haproxy-production.fisdap.svc.tutum.io/"
)

func main() {
	// get /blog => to server1
	// get /login => to server2
	var target1, target2 *url.URL
	var err error

	if target1, err = url.Parse(SERVER1); err != nil {
		log.Fatal("parse url: ", err)
	}

	if target2, err = url.Parse(SERVER2); err != nil {
		log.Fatal("parse url: ", err)
	}

	reverseProxy := new(httputil.ReverseProxy)

	reverseProxy.Director = func(req *http.Request) {
		req.URL.Scheme = "http"

		if strings.HasPrefix(req.URL.Path, "/blog") {
			req.URL.Host = target1.Host
		}

		if strings.HasPrefix(req.URL.Path, "/login") {
			req.URL.Host = target2.Host
		}
	}

	err = http.ListenAndServe(":80", reverseProxy)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}