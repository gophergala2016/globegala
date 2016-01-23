package main

import (
	"log"
	"net/http"

	"github.com/gophergala2016/globegala/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/gophergala2016/globegala/handlers"
)

func main() {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	//	router.GET("/:repo", handlers.Repo)

	log.Fatal(http.ListenAndServe(":8080", router))
}
