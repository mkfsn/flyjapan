package tiger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Do(request *http.Request) (*http.Response, error)
}

type FareCache struct {
	DepartureStation string   `json:"departureStation"`
	ArrivalStation   string   `json:"arrivalStation"`
	Fares            []Fare   `json:"fares"`
	FlightDates      []string `json:"flightDates"`
	UpdatedDate      string   `json:"updatedDate"`
}

type Fare struct {
	FareAmount         float64 `json:"fareAmount"`
	TaxesAndFeesAmount float64 `json:"taxesAndFeesAmount"`
	TotalAmount        float64 `json:"totalAmount"`
	Available          int     `json:"available"`
	DepartureDate      string  `json:"departureDate"`
	Day                int     `json:"day"`
	Ratio              string  `json:"ratio"`
}

func NewFareCacheFromClient(ctx context.Context, client Client, path string) (*FareCache, error) {
	// https://static.tigerairtw.com/fare-cache/TPE:HND:TWD:2020-01
	url := fmt.Sprintf("https://static.tigerairtw.com/fare-cache/%s", path)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return NewFareCacheFromReader(res.Body)
}

func NewFareCacheFromReader(r io.Reader) (*FareCache, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return NewFareCacheFromBytes(body)
}

func NewFareCacheFromBytes(b []byte) (*FareCache, error) {
	var result FareCache
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, fmt.Errorf("tiger: failed to unmarshal: %w", err)
	}
	return &result, nil
}
