package db

import (
	"github.com/lib/pq"
)

const UniqueViolation = "23505"

var ErrUniqueViolation = &pq.Error{
	Code: UniqueViolation,
}
