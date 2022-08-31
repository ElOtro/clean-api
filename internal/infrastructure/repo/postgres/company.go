package repo

import (
	"context"
	"time"

	"github.com/ElOtro/clean-api/internal/entity"
	"github.com/ElOtro/clean-api/pkg/postgres"
)

// CompanyRepo -.
type CompanyRepo struct {
	*postgres.Postgres
}

// New -.
func NewCompanyRepo(pg *postgres.Postgres) *CompanyRepo {
	return &CompanyRepo{pg}
}

// GetHistory -.
func (m CompanyRepo) GetAll() ([]*entity.Company, error) {
	// Construct the SQL query to retrieve all records.
	query := `SELECT id, logo, name, full_name, company_type, details, user_id, created_at, updated_at
		      FROM companies`

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() to execute the query. This returns a sql.Rows resultset
	// containing the result.
	rows, err := m.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	// Importantly, defer a call to rows.Close() to ensure that the resultset is closed
	// before GetAll() returns.
	defer rows.Close()

	companies := []*entity.Company{}

	// Use rows.Next to iterate through the rows in the resultset.
	for rows.Next() {
		// Initialize an empty struct to hold the data for an individual.
		var company entity.Company

		// Scan the values from the row into the struct. Again, note that we're
		// using the pq.Array() adapter on the genres field here.
		err := rows.Scan(
			&company.ID,
			&company.Logo,
			&company.Name,
			&company.FullName,
			&company.CompanyType,
			&company.Details,
			&company.UserID,
			&company.CreatedAt,
			&company.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Add the Company struct to the slice.
		companies = append(companies, &company)
	}

	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error
	// that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}
