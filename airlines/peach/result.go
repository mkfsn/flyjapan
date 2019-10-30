package peach

import (
	"encoding/json"
	"fmt"
	"github.com/mkfsn/flyjapan/airlines"
	"log"
	"regexp"
	"strconv"
	"time"
)

var (
	flightResultsRegexp = regexp.MustCompile("var flightResults = (.*);")
)

func extractFlightResultFromBytes(b []byte) (airlines.Result, error) {
	matches := flightResultsRegexp.FindSubmatch(b)
	if len(matches) != 2 {
		return nil, fmt.Errorf("no matched result (%d)", len(matches))
	}

	var flightResult [2][]Flight
	err := json.Unmarshal(matches[1], &flightResult)
	if err != nil {
		return nil, err
	}

	var inbound, outbound []airlines.Flight
	for _, flight := range flightResult[0] {
		inbound = append(inbound, convertFlightResult(flight))
	}
	for _, flight := range flightResult[1] {
		outbound = append(outbound, convertFlightResult(flight))
	}

	return airlines.GenerateResult(inbound, outbound), nil
}

func convertFlightResult(flight Flight) airlines.Flight {
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
	res := airlines.Flight{
		DepartureTime: departureTime,
		ArrivalTime:   arrivalTime,
		FlightID:      flight.FlightId,
		Origin:        airlines.Airport{Code: flight.OriginCode, Name: flight.Origin},
		Destination:   airlines.Airport{Code: flight.DestinationCode, Name: flight.Destination},
		TaxAdult:      int(taxAdult),
		TaxChild:      flight.TaxChild,
		TaxInfant:     flight.TaxInfant,
		Fares:         make([]*airlines.Fare, 0, len(flight.Fares)),
	}
	for _, fare := range flight.Fares {
		res.Fares = append(res.Fares, &airlines.Fare{
			Seat:     fare.Seat,
			BaseFare: fare.BaseFare,
		})
	}
	return res
}
