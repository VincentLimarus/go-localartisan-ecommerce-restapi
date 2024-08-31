package outputs

import "localArtisans/models/responsesDTO"

type GetAllCartsOutputDTO struct {
	PaginationOutput
	BaseOutput
	Data []responsesDTO.CartResponseDTO `json:"data"`
}

type GetAllCartsByUserIDOutputDTO struct {
	BaseOutput
	Data []responsesDTO.CartResponseDTO `json:"data"`
}

type GetCartOutputDTO struct {
	BaseOutput
	Data responsesDTO.CartResponseDTO `json:"data"`
}

type CreateCartOutputDTO struct {
	BaseOutput
	Data responsesDTO.CartResponseDTO `json:"data"`
}

type DeleteCartOutputDTO struct {
	BaseOutput
	Data responsesDTO.CartResponseDTO `json:"data"`
}

type CheckoutProductFromCartOutputDTO struct {
	BaseOutput
	Data responsesDTO.CartResponseDTO `json:"data"`
}
