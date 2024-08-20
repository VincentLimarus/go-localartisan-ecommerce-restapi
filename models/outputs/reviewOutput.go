package outputs

import "localArtisans/models/responsesDTO"

type GetAllReviewsResponse struct {
	BaseOutput
	PaginationOutput
	Data []responsesDTO.ReviewsResponseDTO `json:"data"`
}

type GetAllReviewsByProductIDResponse struct {
	BaseOutput
	Data []responsesDTO.ReviewsResponseDTO `json:"data"`
}

type GetReviewResponse struct {
	BaseOutput
	Data responsesDTO.ReviewsResponseDTO `json:"data"`
}

type CreateReviewResponse struct {
	BaseOutput
	Data responsesDTO.ReviewsResponseDTO `json:"data"`
}

type DeleteReviewResponse struct {
	BaseOutput
	Data responsesDTO.ReviewsResponseDTO `json:"data"`
}
