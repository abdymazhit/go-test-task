package handler

import (
	"encoding/json"
	"github.com/abdymazhit/go-test-task/model"
	"github.com/valyala/fasthttp"
	"log"
)

func newErrorResponse(ctx *fasthttp.RequestCtx, statusCode int, message string) {
	response, err := json.Marshal(model.ErrorResponse{
		Status: "failure",
		Error:  message,
	})
	if err != nil {
		log.Println("Error sending an erroneous response: " + err.Error())
		return
	}

	ctx.Error(string(response), statusCode)
}
