package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "Address to listen")

func main() {
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Println("Listen on " + *addr)
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
