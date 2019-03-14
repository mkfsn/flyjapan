package flyjapan

import (
	"sort"
)

type Fare struct {
	BaseFare     int    `json:"baseFare"`
	BookingClass string `json:"bookingClass"`
	BookingType  string `json:"bookingType"`
	Discounted   bool   `json:"discounted"`
	Fare         int    `json:"fare"`
	FareCode     string `json:"fareCode"`
	FareId       string `json:"fareId"`
	IsMin        bool   `json:"isMin"`
	IsSale       bool   `json:"isSale"`
	IsStaff      bool   `json:"isStaff"`
	Seat         int    `json:"seat"`
}

type Flight struct {
	ArrivalTime       string          `json:"arrivalTime"`
	ArrivalTimezone   string          `json:"arrivalTimezone"`
	DepartureTime     string          `json:"departureTime"`
	DepartureTimezone int             `json:"departureTimezone"`
	Destination       string          `json:"destination"`
	DestinationCode   string          `json:"destinationCode"`
	Fares             map[string]Fare `json:"fares"`
	FlightDuration    int             `json:"flightDuration"`
	FlightId          string          `json:"flightId"`
	FlightNumber      string          `json:"flightNumber"`
	Origin            string          `json:"origin"`
	OriginCode        string          `json:"originCode"`
	TaxAdult          string          `json:"taxAdult"`
	TaxChild          int             `json:"taxChild"`
	TaxInfant         int             `json:"taxInfant"`
}

func (f Flight) CheapestFare() Fare {
	var key string
	min := -1
	for k, v := range f.Fares {
		if min == -1 || min > v.BaseFare {
			min = v.BaseFare
			key = k
		}
	}
	return f.Fares[key]
}

type Flights []Flight

func (flights Flights) Cheapest() Flight {
	if len(flights) == 0 {
		return Flight{}
	}
	sort.Sort(flights)
	return flights[0]
}

func (f Flights) Len() int {
	return len(f)
}

func (f Flights) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Flights) Less(i, j int) bool {
	return f[i].CheapestFare().BaseFare < f[j].CheapestFare().BaseFare
}
