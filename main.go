package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gophergala2016/globegala/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/gophergala2016/globegala/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	router.GET("/", handlers.Index)
	router.GET("/github/repos", handlers.GetGithubRepos)

	log.Printf("Server Started @ %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
