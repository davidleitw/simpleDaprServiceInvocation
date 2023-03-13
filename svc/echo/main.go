package main

import (
	"echo/function"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()
	r.POST("/", function.Handler)
	if err := fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
		log.Println("addition echo http handler set up error: ", err.Error())
	}
}
