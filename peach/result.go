package peach

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/mkfsn/flyjapan"
)

var (
	flightResultsRegexp = regexp.MustCompile("var flightResults = (.*);")
)

func extractFlightResultFromBytes(b []byte) (flyjapan.Result, error) {
	matches := flightResultsRegexp.FindSubmatch(b)
	if len(matches) != 2 {
		fmt.Printf(string(b))
		return nil, fmt.Errorf("no matched result (%d)", len(matches))
	}

	var flightResult [2][]flyjapan.Flight
	err := json.Unmarshal(matches[1], &flightResult)
	if err != nil {
		return nil, err
	}

	return flyjapan.GenerateResult(flightResult[0], flightResult[1]), nil
}
