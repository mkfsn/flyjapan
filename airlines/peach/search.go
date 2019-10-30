package peach

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"

	"github.com/mkfsn/flyjapan/airlines"
)

func (p *peach) Search(ctx context.Context, q airlines.Query) (airlines.Result, error) {
	b, err := p.doRequest(ctx, q)
	if err != nil {
		return nil, err
	}
	res, err := extractFlightResultFromBytes(b)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *peach) doRequest(ctx context.Context, q airlines.Query) ([]byte, error) {
	req, err := http.NewRequest("POST", "https://booking.flypeach.com/tw", parseQuery(q))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.Header.Set("Origin", "https://www.flypeach.com")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36")
	req.Header.Set("Referer", "https://www.flypeach.com/pc/tw")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// log.Printf("Jar.Cookies: %+v\n---\n", p.client.Jar.Cookies(req.URL))

	// Check if the server sent compressed data
	var reader io.ReadCloser
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(res.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
	default:
		reader = res.Body
	}

	buf := &bytes.Buffer{}
	if _, err := buf.ReadFrom(reader); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
