package outputs

import "localArtisans/models/responsesDTO"

type GetAllUserOutput struct {
	BaseOutput
	PaginationOutput
	Data []responsesDTO.UserResponseDTO `json:"data"`
}

type GetUserOutput struct {
	BaseOutput
	Data responsesDTO.UserResponseDTO `json:"data"`
}

type RegisterUserOutput struct {
	BaseOutput
	Data responsesDTO.UserResponseDTO `json:"data"`
}

type LoginUserOutput struct {
	BaseOutput
	Data responsesDTO.UserResponseDTO `json:"data"`
}

type UpdateUserOutput struct {
	BaseOutput
	Data responsesDTO.UserResponseDTO `json:"data"`
}

type DeleteUserOutput struct {
	BaseOutput
	Data responsesDTO.UserResponseDTO `json:"data"`
}

type ChangePasswordOutput struct {
	BaseOutput
}