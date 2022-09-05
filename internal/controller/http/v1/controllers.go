package v1

import "github.com/ElOtro/clean-api/internal/usecase"

// Create a Models struct which wraps all models.
type Controllers struct {
	Company CompanyController
	Product ProductController
}

// For ease of use, we also add a NewControllers() method which returns a Controllers struct
func NewControllers(usecases *usecase.UseCases) Controllers {
	return Controllers{
		Company: *NewCompanyController(&usecases.Company),
		Product: *NewProductController(&usecases.Product),
	}
}
