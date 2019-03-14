package flyjapan

import (
	"context"
)

type Searcher interface {
	Search(context.Context, Query) (Result, error)
}
