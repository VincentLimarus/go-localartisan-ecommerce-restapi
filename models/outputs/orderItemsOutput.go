package outputs

import "localArtisans/models/responsesDTO"

type GetAllOrderItemsOutput struct {
	BaseOutput
	PaginationOutput
	Data []responsesDTO.OrderItemsResponseDTO `json:"data"`
}

type GetOrderItemByIDOutput struct {
	BaseOutput
	Data responsesDTO.OrderItemsResponseDTO `json:"data"`
}

type GetAllOrderItemsByOrderIDOutput struct {
	BaseOutput
	Data []responsesDTO.OrderItemsResponseDTO `json:"data"`
}

type GetAllOrderItemsByProductIDOutput struct {
	BaseOutput
	Data []responsesDTO.OrderItemsResponseDTO `json:"data"`
}