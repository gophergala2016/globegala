package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/gophergala2016/globegala/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
	"github.com/gophergala2016/globegala/geocoding"
	"github.com/gophergala2016/globegala/github"
)

type AllReposData struct {
	data []RepoData
}

type RepoData struct {
	Name         string
	Contributors []github.Contributor
	Commits      int64
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	uri := r.URL.Path
	if uri == "/" {
		uri = "index.html"
	}
	path := filepath.Join("static", uri)

	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Printf("ioutil.ReadFile('%s') failed with '%s'\n", path, err)
		fmt.Fprint(w, err.Error())
		return
	}

	if len(data) == 0 {
		fmt.Fprint(w, "Asset is empty")
		return
	}

	fmt.Fprint(w, string(data[:len(data)]))
}

func GetGithubRepos(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonRepo, err := ioutil.ReadFile("map_data.json")
	if err != nil {
		allReposData := getDataFromAPI()
		//fmt.Println(allReposData)
		jsonRepo, err = json.MarshalIndent(allReposData.data, "", "\t")
		if err != nil {
			fmt.Println("error:", err)
		}

		if err := ioutil.WriteFile("map_data.json", jsonRepo, 0777); err != nil {
			fmt.Println("writefile err: ", err)
		}
	}

	fmt.Fprint(w, string(jsonRepo[:len(jsonRepo)]))
}

func getDataFromAPI() AllReposData {
	start := time.Now()
	fmt.Printf("Called GetGithubRepos\n")
	allReposData := AllReposData{}

	repos, err := github.FetchAllRepos()
	if err != nil {
		fmt.Printf("Couldn't fetch repos: %v", err.Error())
		return allReposData
	}

	for i := range repos {
		repoData := RepoData{}
		repo := repos[i]
		contributors, err := github.FetchAllContributors(repo.Name)
		if err != nil {
			//			fmt.Printf("FetchAllContributors: %v", err)
		}

		if len(contributors) == 0 {
			continue
		}

		var wg sync.WaitGroup
		wg.Add(len(contributors))
		for i := range contributors {
			contributor := contributors[i]

			go func() {
				defer wg.Done()

				c, err := github.FetchContributor(contributor.Login)
				if err != nil {
					fmt.Printf("FetchContributor: %v", err)
				}

				g, err := geocoding.FetchLatLong(c.Location)
				if err != nil {
					fmt.Printf("FetchLatLong: %v", err)
				}
				c.Geolocation = g

				repoData.Contributors = append(repoData.Contributors, c)
			}()
		}

		wg.Wait()
		repoData.Name = repo.Name

		commits, err := github.FetchRepoCommits(repo.Name)
		if err != nil {
			fmt.Printf("FetchRepoCommits: %v", err)
		}
		repoData.Commits = int64(len(commits))

		allReposData.data = append(allReposData.data, repoData)
	}
	fmt.Printf("Passed Time: %v\n", time.Since(start))

	return allReposData
}
