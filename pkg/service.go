package pkg

import (
	"github.com/tradmark/pkg/search"
	"github.com/tradmark/pkg/tradmark"
)

var (
	TradesRepository tradmark.Repository = nil
	SearchRepository search.Repository   = nil
)

func init() {
	TradesRepository = tradmark.PostgresRepo()
	SearchRepository = search.PostgresRepo()
}
