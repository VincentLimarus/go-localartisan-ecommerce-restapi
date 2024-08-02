package outputs

import "localArtisans/models/responsesDTO"

type GetAllCategoryOutput struct {
	BaseOutput
	PaginationOutput
	Data []responsesDTO.CategoryResponseDTO `json:"data"`
}

type GetCategoryOutput struct {
	BaseOutput
	Data responsesDTO.CategoryResponseDTO `json:"data"`
}

type CreateCategoryOutput struct {
	BaseOutput
	Data responsesDTO.CategoryResponseDTO `json:"data"`
}

type UpdateCategoryOutput struct {
	BaseOutput
	Data responsesDTO.CategoryResponseDTO `json:"data"`
}

type DeleteCategoryOutput struct {
	BaseOutput
	Data responsesDTO.CategoryResponseDTO `json:"data"`
}
