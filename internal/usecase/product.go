package usecase

import (
	"fmt"

	"github.com/ElOtro/clean-api/internal/entity"
)

type ProductRepository interface {
	GetAll() ([]*entity.Product, error)
}

// ProductUseCase -.
type ProductUseCase struct {
	repo ProductRepository
}

// New -.
func NewProductUserCase(r ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repo: r,
	}
}

// History - getting translate history from store.
func (uc *ProductUseCase) List() ([]*entity.Product, error) {
	products, err := uc.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("ProductUseCase - History - s.repo.GetHistory: %w", err)
	}

	return products, nil
}
