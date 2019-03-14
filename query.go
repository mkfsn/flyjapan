package flyjapan

import (
	"time"
)

type Query struct {
	DepartureDate        time.Time
	ReturnDate           time.Time
	DepartureAirportCode string
	ArrivalAirportCode   string
	IsReturn             bool
	AdultCount           int
	ChildCount           int
	InfantCount          int
}
