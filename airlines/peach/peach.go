package peach

import (
	"github.com/mkfsn/flyjapan/airlines"
	"net/http"
	"net/http/cookiejar"
)

type peach struct {
	client *http.Client
}

func New() (airlines.Searcher, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	peach := &peach{
		client: &http.Client{Jar: jar},
	}
	return peach, nil
}
