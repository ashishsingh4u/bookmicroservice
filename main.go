package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ashishsingh4u/bookmicroservice/config"
	"github.com/ashishsingh4u/bookmicroservice/controllers"
	"github.com/ashishsingh4u/bookmicroservice/models"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()

	router := gin.Default()

	pprof.Register(router)

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.CustomRecovery(func(ctx *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			ctx.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}))

	// This middleware is already used when router is created
	// with router := gin.Default(). Use gin.New() to start router with no middleware attached
	// and then use specific middleware like logger or Recovery
	// router.Use(gin.Recovery())

	config := config.GetConfig()
	machineIP := fmt.Sprintf("%s:%s", config.SERVER_IP, config.PORT)
	log.Printf("Server will be starting on %s\n", machineIP)

	router.SetTrustedProxies([]string{config.SERVER_IP})

	router.GET("/", func(ctx *gin.Context) {
		// If the client is 192.168.86.22, use the X-Forwarded-For
		// header to deduce the original client IP from the trust-
		// worthy parts of that header.
		// Otherwise, simply return the direct client IP
		fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
	})

	models.ConnectDatabase()

	// Grouping
	v1 := router.Group("/v1")
	{
		v1.GET("/books", controllers.FindBooks)
		v1.POST("/books", controllers.CreateBook)
		v1.GET("/books/:id", controllers.FindBook)
		v1.PATCH("/books/:id", controllers.UpdateBook)
		v1.DELETE("/books/:id", controllers.DeleteBook)
	}

	router.Run(machineIP)
}
