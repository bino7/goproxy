package main

import (
	"net/http"
	"os"
	"fmt"
	"log"
)

func main() {
	/*proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	log.Fatal(http.ListenAndServe(addr, proxy))*/
	http.HandleFunc("/hi", hi)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	log.Println(port)
	http.ListenAndServe(addr, nil)
}

func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
	log.Println("hi")
}