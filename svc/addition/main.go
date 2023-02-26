package main

import (
	"addition/function"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()
	r.POST("/", function.Handler)
	if err := fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
		log.Println("addition service http handler set up error: ", err.Error())
	}
}
