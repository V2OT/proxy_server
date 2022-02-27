package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "Address to listen")
var verbose = flag.Bool("verbose", false, "Show verbose log")

func main() {
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose
	log.Println("Listen on " + *addr)
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
