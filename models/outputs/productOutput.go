package outputs

import "localArtisans/models/responsesDTO"

type GetAllProductOutput struct {
	BaseOutput
	PaginationOutput
	Data []responsesDTO.ProductResponseDTO `json:"data"`
}

type GetProductOutput struct {
	BaseOutput
	Data responsesDTO.ProductResponseDTO `json:"data"`
}

type CreateProductOutput struct {
	BaseOutput
	Data responsesDTO.ProductResponseDTO `json:"data"`
}

type UpdateProductOutput struct {
	BaseOutput
	Data responsesDTO.ProductResponseDTO `json:"data"`
}

type DeleteProductOutput struct {
	BaseOutput
	Data responsesDTO.ProductResponseDTO `json:"data"`
}