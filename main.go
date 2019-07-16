package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	mux := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			fmt.Fprintf(ctx, "Echo Wep Api :D")
			break
		case "/echo":
			fmt.Fprintf(ctx, "Echo")
			break
		case "/ping":
			fmt.Fprintf(ctx, "pong")
			break
		default:
			ctx.Error("Not Found", fasthttp.StatusNotFound)
			break
		}
	}

	log.Println("Http server listen on 80")
	fasthttp.ListenAndServe(":80", mux)
}
