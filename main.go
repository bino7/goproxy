package main

import (
	"log"
	"net/http"
	"github.com/elazarl/goproxy"
	"os"
	"fmt"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, proxy))
}