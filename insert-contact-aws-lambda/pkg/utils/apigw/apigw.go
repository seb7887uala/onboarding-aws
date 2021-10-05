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

func BadRequestResponse(errMsg string) Response {
	return Response{
		StatusCode: http.StatusBadRequest,
		Body:       errMsg,
	}
}

func InternalErrorResponse() Response {
	return Response{
		StatusCode: http.StatusInternalServerError,
	}
}
