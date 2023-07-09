package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/wilo0087/qrioso-server/internal/repository"
	"github.com/wilo0087/qrioso-server/routes"
)

const defaultPort = ":3000"

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)

	isLocalEnvironment := os.Getenv("_LAMBDA_SERVER_PORT") == "" && os.Getenv("_AWS_LAMBDA_RUNTIME_API") == ""

	r := gin.Default()
	db := repository.NewPostgresConnection()

	routes.RegisterUserRoutes(routes.NewUserRoutes(db), r)

	if !isLocalEnvironment {
		lambda.Start(ginadapter.New(r).ProxyWithContext)
		return
	}

	http.ListenAndServe(defaultPort, r)
}
