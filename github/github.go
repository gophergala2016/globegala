package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var githubAPI = "https://api.github.com"

type Repos []struct {
	Name string `json:"name"`
}

type Contributors []struct {
	Login string
}

func FetchAllRepos() (Repos, error) {
	reqUrl := githubAPI + "/orgs/gophergala2016/repos"

	var repos Repos
	respBody, err := doGetRequest(reqUrl)
	if err != nil {
		return repos, err
	}

	if err := json.Unmarshal(respBody, &repos); err != nil {
		return repos, fmt.Errorf("Unmarshal error: ", err)
	}

	return repos, nil
}

//func FetchAllContributors(repo string) (*Contributors, error) {
//	reqUrl := fmt.Sprintf("%s/repos/gophergala2016/%s/contributors", githubAPI, repo)
//
//	respBody, err := doGetRequest(reqUrl)
//	if err != nil {
//		return nil, err
//	}
//
//	var c *Contributors
//	if err := json.Unmarshal(respBody, c); err != nil {
//		return nil, fmt.Errorf("Unmarshal error: %v", err)
//	}
//
//	return c, nil
//}

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

//func GetContributorLocation() {
//}
