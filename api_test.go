package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FetchUserRepos_WithResults(t *testing.T) {

	a := assert.New(t)

	userName := "dring1"
	repos, _ := FetchUserRepoList(userName)
	a.NotEmpty(repos, "Should not be empty")
}

func Test_FetchUserRepos_NoResults(t *testing.T) {
	a := assert.New(t)
	fakeUser := "asfbusbgudsbdjhkzbdyug"
	repos, err := FetchUserRepoList(fakeUser)
	a.Empty(repos, "Should be empty")
	a.Error(err, "Should have errored")
	a.Equal(fmt.Sprintf("No results for user: %s", fakeUser), err.Error())
}

func FakeServer(b string, f func()) {

	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, b)
	}))
	defer ts.Close()

	f()

}
