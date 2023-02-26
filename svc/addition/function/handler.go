package function

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/valyala/fasthttp"
)

type additionApiRequestBody struct {
	N1 int `json:"n1"`
	N2 int `json:"n2"`
}

type additionApiResponseBody struct {
	Result int `json:"result"`
}

func Handler(ctx *fasthttp.RequestCtx) {
	reqBody := &additionApiRequestBody{}

	if err := json.Unmarshal(ctx.PostBody(), &reqBody); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("addition event: n1 = %d, n2 = %d\n", reqBody.N1, reqBody.N2)

	resBody := &additionApiResponseBody{Result: reqBody.N1 + reqBody.N2}
	body, err := json.Marshal(&resBody)
	if err != nil {
		log.Println("addition handler get error: ", err.Error())
		ctx.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.Response.SetBody(body)
	ctx.SetStatusCode(http.StatusOK)
}
