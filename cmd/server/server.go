package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/wilo0087/qrioso-server/internal/handlers"
)

func main() {

	log.SetFlags(log.Flags() | log.Lshortfile)

	isLocalEnvironment := os.Getenv("_LAMBDA_SERVER_PORT") == "" && os.Getenv("_AWS_LAMBDA_RUNTIME_API") == ""

	handlers := handlers.StartApiRoutes(handlers.NewDefaultHandler())

	if !isLocalEnvironment {
		lambda.Start(ginadapter.New(handlers).ProxyWithContext)
		return
	}

	// server.NewServer()
	http.ListenAndServe(":3000", handlers)

}
