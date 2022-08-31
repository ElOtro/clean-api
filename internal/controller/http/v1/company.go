package v1

import (
	"net/http"

	"github.com/ElOtro/clean-api/internal/entity"
)

type CompanyUseCase interface {
	List() ([]*entity.Company, error)
}

type CompanyController struct {
	uc CompanyUseCase
}

func NewCompanyController(uc CompanyUseCase) *CompanyController {
	return &CompanyController{uc: uc}
}

type listCompanyResponse struct {
	Company []*entity.Company `json:"companies"`
}

// @Summary     Show company list
// @Description Show all company list
// @ID          companyList
// @Tags        companies
// @Accept      json
// @Produce     json
// @Success     200 {object} listCompanyResponse
// @Router      /companies [get]
func (c *CompanyController) List(w http.ResponseWriter, r *http.Request) {
	companies, err := c.uc.List()
	if err != nil {
		errorResponse(w, r, http.StatusInternalServerError, "get companies")

		return
	}

	err = writeJSON(w, http.StatusOK, listCompanyResponse{companies}, nil)
	if err != nil {
		serverErrorResponse(w, r, err)
	}
}
