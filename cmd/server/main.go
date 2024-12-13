package main

import (
	"net/http"

	_ "github.com/LeVanHieu0509/backend-go/cmd/swag/docs"
	"github.com/LeVanHieu0509/backend-go/internal/initialize"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count_total",
		Help: "Total number of ping requests.",
	},
)

func ping(c *gin.Context) {
	pingCounter.Inc()
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

// swagger embed files
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server caller server.
// @termsOfService  github.com/LeVanHieu0509/go-backend-api

// @contact.name   TEAM TIP GO
// @contact.url    github.com/LeVanHieu0509/go-backend-api
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8001
// @BasePath  /v1/2024

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {

	r := initialize.Run()

	prometheus.MustRegister(pingCounter)

	r.GET("/ping/200", ping)
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8001")
}
