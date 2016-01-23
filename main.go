package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gophergala2016/globegala/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/gophergala2016/globegala/handlers"
)

func main() {
	router := httprouter.New()
	router.GET("/", handlers.Index)
	router.GET("/test", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "TESTTT")
	})

	port := 5151
	log.Printf("Server Started @ %d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
