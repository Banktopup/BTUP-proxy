package main

import (
	"gopkg.in/elazarl/goproxy.v1"
	"log"
	"net/http"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8080", proxy))

}