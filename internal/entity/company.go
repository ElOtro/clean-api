package entity

import "time"

// CompanyDetails type details
type CompanyDetails struct {
	INN     string `json:"inn,omitempty"`
	KPP     string `json:"kpp,omitempty"`
	OGRN    string `json:"ogrn,omitempty"`
	Address string `json:"address,omitempty"`
}

// Company type
type Company struct {
	ID          int64           `json:"id"`
	Logo        *string         `json:"logo,omitempty"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name,omitempty"`
	CompanyType int             `json:"company_type,omitempty"`
	Details     *CompanyDetails `json:"details,omitempty"`
	DestroyedAt *time.Time      `json:"destroyed_at,omitempty"`
	CreatedAt   *time.Time      `json:"created_at,omitempty"`
	UpdatedAt   *time.Time      `json:"updated_at,omitempty"`
}

// CompanySearch  type
type CompanySearch struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CompanyFilters struct {
	Name string
}
