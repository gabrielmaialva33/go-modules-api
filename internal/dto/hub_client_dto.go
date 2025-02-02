package dto

type CreateHubClientDTO struct {
	Name       string `json:"name" validate:"required,min=3,max=100"`
	ExternalID uint   `json:"external_id" validate:"required"`
}

type UpdateHubClientDTO struct {
	Name       string `json:"name" validate:"omitempty,min=3,max=100"`
	ExternalID uint   `json:"external_id" validate:"omitempty"`
}

type GetHubClientDTO struct {
	Search    string `json:"search" validate:"omitempty"`
	Active    *bool  `json:"active" validate:"omitempty"`
	SortField string `json:"sort_field" validate:"omitempty,oneof=id name active external_id created_at updated_at"`
	SortOrder string `json:"sort_order" validate:"omitempty,oneof=asc desc"`
}

type PaginatedHubClientDTO struct {
	Page      int    `json:"page" validate:"omitempty,min=1"`
	PageSize  int    `json:"page_size" validate:"omitempty,min=1,max=100"`
	Search    string `json:"search" validate:"omitempty"`
	Active    *bool  `json:"active" validate:"omitempty"`
	SortField string `json:"sort_field" validate:"omitempty,oneof=id name active external_id created_at updated_at"`
	SortOrder string `json:"sort_order" validate:"omitempty,oneof=asc desc"`
}
