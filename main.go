package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "index")
}

func Repo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s", ps.ByName("repo"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/:repo", Repo)

	log.Fatal(http.ListenAndServe(":8080", router))
}
