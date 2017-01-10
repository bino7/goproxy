package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	port := 8080
	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
