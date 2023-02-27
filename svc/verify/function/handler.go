package function

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/valyala/fasthttp"
)

const (
	sidecarAddress       = "localhost:3500"
	enableSidecarAddress = true
)

var proxyClient = &fasthttp.HostClient{
	Addr: sidecarAddress,
}

type verifyApiRequestBody struct {
	N1 int `json:"n1"`
	N2 int `json:"n2"`
}

func proxyHandler(ctx *fasthttp.RequestCtx, appId string) {
	log.Println("proxy handler: send request to addition service")
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.Set("dapr-app-id", appId)
	req.Header.SetContentType("application/json")
	if enableSidecarAddress {
		log.Println("enable sidecar address with req.SetRequestURI")
		req.SetRequestURI("http://localhost:3500")
	}

	req.SetBodyRaw(ctx.PostBody())
	if err := proxyClient.Do(req, res); err != nil {
		log.Println("proxy handler, get error: ", err.Error())
	}

	log.Printf("proxy handler got result, status code = %d, body = %s", res.StatusCode(), string(res.Body()))

	ctx.Response.SetBodyRaw(res.Body())
	ctx.SetStatusCode(res.StatusCode())
}

func Handler(ctx *fasthttp.RequestCtx) {
	reqBody := &verifyApiRequestBody{}

	if err := json.Unmarshal(ctx.PostBody(), &reqBody); err != nil {
		ctx.Error(err.Error(), http.StatusBadRequest)
		return
	}

	if reqBody.N1 < 0 || reqBody.N2 < 0 {
		ctx.Error("n1, n2 must bigger than zero", http.StatusBadRequest)
		return
	}

	log.Printf("Verify event, n1 = %d, n2 = %d, verified.\n", reqBody.N1, reqBody.N2)
	proxyHandler(ctx, "addition")
}
