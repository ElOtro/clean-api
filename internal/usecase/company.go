package usecase

import (
	"fmt"

	"github.com/ElOtro/clean-api/internal/entity"
)

type CompanyRepository interface {
	GetAll() ([]*entity.Company, error)
}

// CompanyUseCase -.
type CompanyUseCase struct {
	repo CompanyRepository
}

// New -.
func NewCompanyUseCase(r CompanyRepository) *CompanyUseCase {
	return &CompanyUseCase{
		repo: r,
	}
}

// History - getting translate history from store.
func (uc *CompanyUseCase) List() ([]*entity.Company, error) {
	companies, err := uc.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("CompanyUseCase - History - s.repo.GetHistory: %w", err)
	}

	return companies, nil
}
