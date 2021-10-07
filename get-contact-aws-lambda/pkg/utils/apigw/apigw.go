package apigw

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type Request = events.APIGatewayProxyRequest
type Response = events.APIGatewayProxyResponse

func OkResponse(body string) Response {
	return Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}
}

func NotFoundResponse(msg string) Response {
	return Response{
		StatusCode: http.StatusNotFound,
		Body:       msg,
	}
}

func InternalErrResponse() Response {
	return Response{
		StatusCode: http.StatusInternalServerError,
	}
}
