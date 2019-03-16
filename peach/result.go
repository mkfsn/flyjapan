package peach

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

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

	var flightResult [2][]Flight
	err := json.Unmarshal(matches[1], &flightResult)
	if err != nil {
		return nil, err
	}

	var inbound, outbound []flyjapan.Flight
	for _, flight := range flightResult[0] {
		inbound = append(inbound, convertFlightResult(flight))
	}
	for _, flight := range flightResult[1] {
		outbound = append(outbound, convertFlightResult(flight))
	}

	return flyjapan.GenerateResult(inbound, outbound), nil
}

func convertFlightResult(flight Flight) flyjapan.Flight {
	format := "2006/01/02 15:04:05" // 2019/03/22 21:00:00
	departureTime, err := time.Parse(format, flight.DepartureTime)
	if err != nil {
		log.Printf("Failed to parse time %s: %s\n", flight.DepartureTime, err.Error())
	}
	arrivalTime, err := time.Parse(format, flight.ArrivalTime)
	if err != nil {
		log.Printf("Failed to parse time %s: %s\n", flight.DepartureTime, err.Error())
	}
	taxAdult, _ := strconv.ParseInt(flight.TaxAdult, 10, 64)
	res := flyjapan.Flight{
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
		FlightID:      flight.FlightId,
		Origin:        flyjapan.Airport{Code: flight.OriginCode, Name: flight.Origin},
		Destination:   flyjapan.Airport{Code: flight.DestinationCode, Name: flight.Destination},
		TaxAdult:      int(taxAdult),
		TaxChild:      flight.TaxChild,
		TaxInfant:     flight.TaxInfant,
		Fares:         make([]*flyjapan.Fare, 0, len(flight.Fares)),
	}
	for _, fare := range flight.Fares {
		res.Fares = append(res.Fares, &flyjapan.Fare{
			Seat:     fare.Seat,
			BaseFare: fare.BaseFare,
		})
	}
	return res
}
