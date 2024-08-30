package outputs

import "localArtisans/models/responsesDTO"

type GetAllReviewsOutput struct {
	BaseOutput
	PaginationOutput
	Data []responsesDTO.ReviewsResponseDTO `json:"data"`
}

type GetAllReviewsByProductIDOutput struct {
	BaseOutput
	Data []responsesDTO.ReviewsResponseDTO `json:"data"`
}

type GetReviewOutput struct {
	BaseOutput
	Data responsesDTO.ReviewsResponseDTO `json:"data"`
}

type CreateReviewOutput struct {
	BaseOutput
	Data responsesDTO.ReviewsResponseDTO `json:"data"`
}

type DeleteReviewOutput struct {
	BaseOutput
	Data responsesDTO.ReviewsResponseDTO `json:"data"`
}
