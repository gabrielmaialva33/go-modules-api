package dto

type CreateRoleDTO struct {
	Name string `json:"name" validate:"required,min=3,max=50"`
	Slug string `json:"slug" validate:"required,min=3,max=50"`
}

type UpdateRoleDTO struct {
	Name   string `json:"name" validate:"omitempty,min=3,max=50"`
	Slug   string `json:"slug" validate:"omitempty,min=3,max=50"`
	Active *bool  `json:"active" validate:"omitempty"`
}

type ListRoleDTO struct {
	Search    string `json:"search" validate:"omitempty"`
	Active    *bool  `json:"active" validate:"omitempty"`
	SortField string `json:"sort_field" validate:"omitempty,oneof=id name slug active created_at updated_at"`
	SortOrder string `json:"sort_order" validate:"omitempty,oneof=asc desc"`
}

type PaginatedRoleDTO struct {
	Page      int    `json:"page" validate:"omitempty,min=1"`
	PageSize  int    `json:"page_size" validate:"omitempty,min=1,max=100"`
	Search    string `json:"search" validate:"omitempty"`
	Active    *bool  `json:"active" validate:"omitempty"`
	SortField string `json:"sort_field" validate:"omitempty,oneof=id name slug active created_at updated_at"`
	SortOrder string `json:"sort_order" validate:"omitempty,oneof=asc desc"`
}
