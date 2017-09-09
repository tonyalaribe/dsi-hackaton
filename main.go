package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/tonyalaribe/dsi-hackaton/resources"
)

func GetMainEngine() *gin.Engine {

	router := gin.Default()
	//router.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))
	router.Use(cors.New(cors.Config{
		//AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8000", "http://localhost:8282"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "HEAD", "OPTIONS", "DELETE"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
		AllowAllOrigins:  true,
	}))

	router.Use(static.Serve("/", static.LocalFile("./admin", false)))

	router.GET("/ping", func(c *gin.Context) {
		c.Status(200)
	})

	api := router.Group("/api")
	resources.Register("locations", resources.Location{}, api)

	router.NoRoute(func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "./admin/index.html")
		c.Abort()

	})

	return router
}

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Println("No Global port has been defined, using default")
		PORT = "8080"

	}

	go func() {
		GetMainEngine().Run(":8181")
	}()

	GetMainEngine().RunTLS(":"+PORT, "./server.crt", "./server.key")

	//GetMainEngine().Run(":" + PORT)
}
