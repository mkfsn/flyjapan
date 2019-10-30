package airlines

import (
	"time"
)

type Fare struct {
	Seat     int
	BaseFare int
}

type Airport struct {
	Code string // e.g. TPE, KIX, HND
	Name string
}

type Flight struct {
	DepartureTime time.Time
	ArrivalTime   time.Time
	FlightID      string
	Origin        Airport
	Destination   Airport
	Fares         []*Fare // TWD
	TaxAdult      int
	TaxChild      int
	TaxInfant     int
}

func (f *Flight) Cheapest() int {
	if len(f.Fares) == 0 {
		return 0
	}
	min := f.Fares[0].BaseFare
	for i := 1; i < len(f.Fares); i++ {
		if f.Fares[i].BaseFare < min {
			min = f.Fares[i].BaseFare
		}
	}
	return min
}

type Flights []Flight

func (flights Flights) FilterBy(filter FilterFunc) Flights {
	var res Flights
	for _, f := range flights {
		if filter(f) {
			res = append(res, f)
		}
	}
	return res
}
