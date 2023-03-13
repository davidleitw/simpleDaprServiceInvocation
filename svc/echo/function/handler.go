package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

type Message struct {
	Msg string `json:"message"`
}

func Handler(ctx *fasthttp.RequestCtx) {
	body := &Message{}

	if err := json.Unmarshal(ctx.PostBody(), &body); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	message := fmt.Sprintf("receive message %s at %s", body.Msg, time.Now().String())
	log.Println(message)
	ctx.SetBodyString(message)
	ctx.SetStatusCode(http.StatusOK)
}
