package outputs

import "localArtisans/models/responsesDTO"

type GetAllArtisanOutput struct {
	BaseOutput
	PaginationOutput
	Data []responsesDTO.ArtisansResponseDTO `json:"data"`
}

type GetArtisanOutput struct {
	BaseOutput
	Data responsesDTO.ArtisansResponseDTO `json:"data"`
}

type RegisterArtisanOutput struct {
	BaseOutput
	Data responsesDTO.ArtisansResponseDTO `json:"data"`
}

type UpdateArtisanOutput struct {
	BaseOutput
	Data responsesDTO.ArtisansResponseDTO `json:"data"`
}

type DeleteArtisanOutput struct {
	BaseOutput
	Data responsesDTO.ArtisansResponseDTO `json:"data"`
}

