package outputs

import "localArtisans/models/responsesDTO"

type GetAllOrderOutput struct {
	PaginationOutput
	BaseOutput
	Data []responsesDTO.OrderResponseDTO `json:"data"`
}

type GetAllOrderByUserIDOutput struct {
	BaseOutput
	Data []responsesDTO.OrderResponseDTO `json:"data"`
}

type GetOrderOutput struct {
	BaseOutput
	Data responsesDTO.OrderResponseDTO `json:"data"`
}

type CreateOrderOutput struct {
	BaseOutput
	Data responsesDTO.OrderResponseDTO `json:"data"`
}

type DeleteOrderOutput struct {
	BaseOutput
	Data responsesDTO.OrderResponseDTO `json:"data"`
}

type CheckOutOrderOutput struct {
	BaseOutput
	Data responsesDTO.OrderResponseDTO `json:"data"`
}