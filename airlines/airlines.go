package airlines

import (
	"context"
)

type AirlineName string

const (
	AirlinePeach   AirlineName = "peach"
	AirlineTiger               = "tiger"
	AirlineJetstar             = "jetstar"
)

type Airline interface {
	Search(context.Context, Query) (Result, error)
}
