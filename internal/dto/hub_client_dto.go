package dto

type CreateHubClientDTO struct {
	Name       string `json:"name" validate:"required,min=3,max=100"`
	ExternalID uint   `json:"external_id" validate:"required"`
}

type UpdateHubClientDTO struct {
	Name       string `json:"name" validate:"omitempty,min=3,max=100"`
	ExternalID uint   `json:"external_id" validate:"omitempty"`
}
