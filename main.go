package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gophergala2016/globegala/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/gophergala2016/globegala/github"
	"github.com/gophergala2016/globegala/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	github.CacheInit()

	router := httprouter.New()

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	router.GET("/", handlers.Index)
	router.GET("/github/repos", handlers.GetGithubRepos)

	router.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Test")
	})

	log.Printf("Server Started @ %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
