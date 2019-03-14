package peach

import (
	"context"
	"net/http"
	"net/http/cookiejar"

	"github.com/mkfsn/flyjapan"
)

type peach struct {
	client *http.Client
}

func New() (flyjapan.Searcher, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	peach := &peach{
		client: &http.Client{Jar: jar},
	}
	return peach, nil
}

func (p *peach) Search(ctx context.Context, q flyjapan.Query) (flyjapan.Result, error) {
	b, err := p.doRequest(ctx, q)
	if err != nil {
		return nil, err
	}
	return extractFlightResultFromBytes(b)
}
