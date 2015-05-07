package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	BASE = "https://api.github.com"
)

type Repositories struct {
	CreatedAt   *time.Time `json:"created_at"`
	Description string     `json:"description"`
	Size        int        `json:"size"`
	UpdatedAt   *time.Time `json:"created_at"`
	Url         string     `json:"html_url"`
	Name        string     `json:"name"`
}

// type UserRepoSearchResults struct {
// 	repos Repositories
// }
//

func FetchUserRepoList(userName string) ([]string, error) {
	url := fmt.Sprintf("%s/users/%s/repos", BASE, url.QueryEscape(userName))
	results := []Repositories{}
	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, errors.New(fmt.Sprintf("No results for user: %s", userName))
	}

	names := make([]string, 10)

	if len(results) == 0 {
		return names, errors.New(fmt.Sprintf("No results for user: %s", userName))
	}
	for _, value := range results {
		names = append(names, value.Name)
	}

	// body, err := ioutil.ReadAll(res.Body)
	// fmt.Print(os.Stdout, string(body))
	return names, err
}
