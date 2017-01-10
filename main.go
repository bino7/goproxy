package main

import (
	"log"
	"net/http"
	"github.com/elazarl/goproxy"
	"os"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":80", proxy))
}