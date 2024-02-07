package sqlc

import (
	"github.com/cbugk/iamfeelingcody/src/internal/sqlc/sqlite"
)

var q *sqlite.Queries

func Q() *sqlite.Queries {
	return q
}
