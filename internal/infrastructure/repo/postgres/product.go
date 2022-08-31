package repo

import (
	"context"
	"time"

	"github.com/ElOtro/clean-api/internal/entity"
	"github.com/ElOtro/clean-api/pkg/postgres"
)

// ProductRepo -.
type ProductRepo struct {
	*postgres.Postgres
}

// New -.
func NewProductRepo(pg *postgres.Postgres) *ProductRepo {
	return &ProductRepo{pg}
}

// GetHistory -.
func (r *ProductRepo) GetAll() ([]*entity.Product, error) {
	// Construct the SQL query to retrieve all records.
	query := `SELECT id, product_type, name, description, sku, price, created_at, updated_at 
			  FROM products 
			  WHERE destroyed_at IS NULL`

	// Create a context with a 3-second timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Use QueryContext() to execute the query. This returns a sql.Rows resultset
	// containing the result.
	rows, err := r.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	// Importantly, defer a call to rows.Close() to ensure that the resultset is closed
	// before GetAll() returns.
	defer rows.Close()

	products := []*entity.Product{}

	// Use rows.Next to iterate through the rows in the resultset.
	for rows.Next() {
		// Initialize an empty struct to hold the data for an individual record.
		var product entity.Product

		// Scan the values from the row into the struct. Again, note that we're
		// using the pq.Array() adapter on the genres field here.
		err := rows.Scan(
			&product.ID,
			&product.ProductType,
			&product.Name,
			&product.Description,
			&product.SKU,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Add the Product struct to the slice.
		products = append(products, &product)
	}

	// When the rows.Next() loop has finished, call rows.Err() to retrieve any error
	// that was encountered during the iteration.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
