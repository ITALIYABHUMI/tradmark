package pkg

import (
	"github.com/tradmark/pkg/search"
	"github.com/tradmark/pkg/tradmark"
)

var (
	TradesRepository tradmark.Repository
	SearchRepository search.Repository
)

func init() {
	TradesRepository = tradmark.PostgresRepo()
	SearchRepository = search.PostgresRepo()
}
