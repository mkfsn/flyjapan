package airlines

type Result interface {
	InBound() Flights
	OutBound() Flights
}

type result struct {
	inBound  Flights
	outBound Flights
}

func (r *result) InBound() Flights {
	return r.inBound
}

func (r *result) OutBound() Flights {
	return r.outBound
}

func GenerateResult(inBound, outBound Flights) Result {
	return &result{inBound: inBound, outBound: outBound}
}
