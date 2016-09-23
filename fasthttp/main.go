package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/valyala/fasthttp"
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	// ctx.SetContentType("text/plain; charset=utf8")
	ctx.Response.Header.Set("Last-Modified", "Thu, 18 Jun 2015 10:24:27 GMT")
	ctx.Response.Header.Set("Accept-Ranges", "bytes")
	ctx.Response.Header.Set("E-Tag", "55829c5b-17")
	ctx.Response.Header.Set("Server", "golang-http-server")
	fmt.Fprintf(ctx, "%s", "<h1>\nHello world!\n</h1>\n")
}

func main() {
	isAutoProcs := flag.Bool("auto-procs", true, "设置GOMAXPROCS到CPU核数")
	flag.Parse()
	if !*isAutoProcs {
		runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	}

	log.Printf("Go http Server listen on :9501")
	log.Fatal(fasthttp.ListenAndServe(":9501", requestHandler))
}
