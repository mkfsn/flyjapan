package flyjapan

import "github.com/mkfsn/flyjapan/airlines"

// FIXME(mkfsn)
type Result struct {
	Inbound  airlines.Flights
	Outbound airlines.Flights
}
