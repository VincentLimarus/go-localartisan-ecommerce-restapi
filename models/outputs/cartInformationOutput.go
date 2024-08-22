package outputs

import "localArtisans/models/responsesDTO"

type GetAllCartInformationsOutput struct {
	PaginationOutput
	BaseOutput
	Data []responsesDTO.CartInformationResponseDTO `json:"data"`
}

type GetALlCartInformationByCartIDOutput struct {
	BaseOutput
	Data []responsesDTO.CartInformationResponseDTO `json:"data"`
}

type AddItemToCartOutput struct {
	BaseOutput
	Data responsesDTO.CartInformationResponseDTO `json:"data"`
}

type UpdateItemInCartOutput struct {
	BaseOutput
	Data responsesDTO.CartInformationResponseDTO `json:"data"`
}

type DeleteItemInCartOutput struct {
	BaseOutput
	Data responsesDTO.CartInformationResponseDTO `json:"data"`
}

	
