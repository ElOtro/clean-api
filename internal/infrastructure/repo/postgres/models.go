package repo

import (
	"errors"

	"github.com/ElOtro/clean-api/pkg/postgres"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Create a Models struct which wraps all models.
type Models struct {
	Companies CompanyRepo
	Products  ProductRepo
}

// For ease of use, we also add a New() method which returns a Models struct containing
// the initialized InvoiceModel.
func NewModels(pg *postgres.Postgres) Models {
	return Models{
		Companies: CompanyRepo{pg},
		Products:  ProductRepo{pg},
	}
}
