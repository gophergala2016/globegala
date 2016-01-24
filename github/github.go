package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gophergala2016/globegala/geocoding"
)

var (
	githubAPI   = "https://api.github.com"
	accessToken = os.Getenv("access_token")
)

type Repos []struct {
	Name string `json:"name"`
}

type Contributors []struct {
	Login string `json:"login"`
}

type Contributor struct {
	Login    string `json:"login"`
	Location string `json:"location"`

	Geolocation geocoding.Geolocation
}

func FetchAllRepos() (Repos, error) {
	var allRepos Repos
	pageNum := 1

	for i := 0; i < 7; i++ {
		var repos Repos
		reqUrl := fmt.Sprintf("%s/orgs/gophergala2016/repos?access_token=%s&page=%v", githubAPI, accessToken, pageNum)

		respBody, err := doGetRequest(reqUrl)
		if err != nil {
			return repos, err
		}

		fmt.Printf("%v", string(respBody))

		if err := json.Unmarshal(respBody, &repos); err != nil {
			return repos, fmt.Errorf("Unmarshal error: ", err)
		}

		allRepos = append(allRepos, repos...)

		pageNum += 1
	}

	return allRepos, nil
}

func FetchAllContributors(repo string) (Contributors, error) {
	reqUrl := fmt.Sprintf("%s/repos/gophergala2016/%s/contributors?access_token=%s", githubAPI, repo, accessToken)

	respBody, err := doGetRequest(reqUrl)
	if err != nil {
		return nil, err
	}

	var c Contributors
	if err := json.Unmarshal(respBody, &c); err != nil {
		return nil, fmt.Errorf("Unmarshal error: %v", err)
	}

	return c, nil
}

func FetchContributor(contributor string) (Contributor, error) {
	reqUrl := fmt.Sprintf("%s/users/%s?access_token=%s", githubAPI, contributor, accessToken)

	var c Contributor
	respBody, err := doGetRequest(reqUrl)
	if err != nil {
		return c, err
	}

	if err := json.Unmarshal(respBody, &c); err != nil {
		return c, fmt.Errorf("Unmarshal error: %v", err)
	}

	return c, nil
}

func doGetRequest(reqUrl string) ([]byte, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest error: %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll error: %v", err)
	}

	return body, nil
}
