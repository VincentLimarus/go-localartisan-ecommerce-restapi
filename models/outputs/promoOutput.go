package outputs

import "localArtisans/models/responsesDTO"

type GetAllPromoOutput struct {
	BaseOutput
	PaginationOutput
	Data []responsesDTO.PromoResponseDTO `json:"data"`
}

type GetPromoOutput struct {
	BaseOutput
	Data responsesDTO.PromoResponseDTO `json:"data"`
}

type GetAllPromoByProductID struct {
	BaseOutput
	Data []responsesDTO.PromoResponseDTO `json:"data"`
}

type CreatePromoOutput struct {
	BaseOutput
	Data responsesDTO.PromoResponseDTO `json:"data"`
}

type UpdatePromoOutput struct {
	BaseOutput
	Data responsesDTO.PromoResponseDTO `json:"data"`
}

type DeletePromoOutput struct {
	BaseOutput
	Data responsesDTO.PromoResponseDTO `json:"data"`
}
