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

type GetAllProductByArtisanIDOutput struct {
	BaseOutput
	Data []responsesDTO.ProductResponseDTO `json:"data"`
}

type GetAllProductByCategoryIDOutput struct {
	BaseOutput
	Data []responsesDTO.ProductResponseDTO `json:"data"`
}

type AddProductToCartOutput struct {
	BaseOutput
	Data responsesDTO.CartResponseDTO `json:"data"`
}

type CheckoutProductOutput struct {
	BaseOutput
	Data responsesDTO.OrderResponseDTO `json:"data"`
}
