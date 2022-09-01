package usecase

import repo "github.com/ElOtro/clean-api/internal/infrastructure/repo/postgres"

// Create a UseCases struct which wraps all models.
type UseCases struct {
	Company CompanyUseCase
	Product ProductUseCase
}

// For ease of use, we also add a NewUseCases() method which returns a UseCases struct containing
// the initialized UseCases.
func NewUseCases(models *repo.Repo) UseCases {
	return UseCases{
		Company: *NewCompanyUseCase(&models.Companies),
		Product: *NewProductUserCase(&models.Products),
	}
}
