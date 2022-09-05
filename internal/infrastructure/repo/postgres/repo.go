package repo

import (
	"errors"

	"github.com/ElOtro/clean-api/pkg/postgres"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Repo struct which wraps all repo.
type Repo struct {
	Companies CompanyRepo
	Products  ProductRepo
}

// For ease of use, we also add a NewRepo() method which returns a Repo struct
func NewRepo(pg *postgres.Postgres) Repo {
	return Repo{
		Companies: CompanyRepo{pg},
		Products:  ProductRepo{pg},
	}
}
