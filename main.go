package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	fmt.Println(addr)
	log.Println(addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}

