package v1

import (
	"net/http"

	"github.com/ElOtro/clean-api/internal/entity"
)

type ProductUseCase interface {
	List() ([]*entity.Product, error)
}

type ProductController struct {
	uc ProductUseCase
}

func NewProductController(uc ProductUseCase) *ProductController {
	return &ProductController{uc: uc}
}

type listProductResponse struct {
	Product []*entity.Product `json:"products"`
}

// List         godoc
// @Summary     Show product list
// @Description Show all product list
// @ID          productList
// @Tags        products
// @Accept      json
// @Produce     json
// @Success     200 {object} listProductResponse
// @Router      /products [get]
func (c *ProductController) List(w http.ResponseWriter, r *http.Request) {
	products, err := c.uc.List()
	if err != nil {
		errorResponse(w, r, http.StatusInternalServerError, "get products")

		return
	}

	err = writeJSON(w, http.StatusOK, listProductResponse{products}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}
