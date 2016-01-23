package handlers

import (
	"fmt"
	//	"log"
	"net/http"

	"github.com/gophergala2016/globegala/github"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	repos, err := github.FetchAllRepos()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Fprint(w, repos[0].Name)
}

//func Repo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	repo := ps.ByName("repo")
//
//	contributors, err := github.FetchAllContributors(repo)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Fprintf(w, "hello %s", &contributors.Contributor[0].Login)
//}
