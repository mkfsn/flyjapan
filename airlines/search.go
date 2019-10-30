package airlines

import (
	"context"
)

type Searcher interface {
	Search(context.Context, Query) (Result, error)
}
