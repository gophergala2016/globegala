package main

import (
	"log"
	"net/http"

	"github.com/gophergala2016/globegala/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	//	router.GET("/:repo", handlers.Repo)

	log.Fatal(http.ListenAndServe(":8080", router))
}
