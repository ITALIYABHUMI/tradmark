package pkg

import (
	"github.com/tradmark/pkg/tradmark"
)

var (
	TradesRepository tradmark.Repository
)

func init() {
	TradesRepository = tradmark.PostgresRepo()
}
