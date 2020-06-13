package main

import (
	"github.com/lilahamstern/hamsterapps.net/server/internal/auth"
	. "github.com/lilahamstern/hamsterapps.net/server/internal/handler"
	database "github.com/lilahamstern/hamsterapps.net/server/internal/pkg/db/postgres"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := gin.Default()

	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()

	router.POST("/query", GraphqlHandler())
	router.GET("/", PlaygroundHandler())

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
