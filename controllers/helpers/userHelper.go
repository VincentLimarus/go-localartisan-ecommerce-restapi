package helpers

import (
	"context"
	"fmt"
	"localArtisans/configs"
	"localArtisans/models/database"
	outputs "localArtisans/models/outputs"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"
	"localArtisans/utils"

	"github.com/google/uuid"
)

func GetAllUser(GetAllUsersRequestDTO requestsDTO.GetAllUsersRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var users []database.User

	if GetAllUsersRequestDTO.Limit == 0 {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit cannot be 0",
		}
		return 400, output
	}
	
	if GetAllUsersRequestDTO.Limit > 100 {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit cannot be more than 100",
		}
		return 400, output
	}

	offset := (GetAllUsersRequestDTO.Page - 1) * GetAllUsersRequestDTO.Limit
	order := fmt.Sprintf("%s %s", GetAllUsersRequestDTO.OrderBy, GetAllUsersRequestDTO.OrderType)
	err := db.Offset(offset).Limit(GetAllUsersRequestDTO.Limit).Order(order).Find(&users).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	if len(users) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: No data found",
		}
		return 404, output
	}

	var totalData int64
	var totalPage int 
	db.Model(&database.User{}).Count(&totalData)

	if totalData % int64(GetAllUsersRequestDTO.Limit) == 0 {
		totalPage = int(totalData) / GetAllUsersRequestDTO.Limit
	} else {
		totalPage = int(totalData) / GetAllUsersRequestDTO.Limit + 1
	}

	output := outputs.GetAllUserOutput{}
	output.Page = GetAllUsersRequestDTO.Page
	output.Limit = GetAllUsersRequestDTO.Limit
	output.OrderBy = GetAllUsersRequestDTO.OrderBy
	output.OrderType = GetAllUsersRequestDTO.OrderType
	output.Code = 200 
	output.Message = "Success: Data found"
	output.TotalData = int(totalData)
	output.TotalTake = len(users)
	output.TotalPage = totalPage
	
	for _, user := range users {
		output.Data = append(output.Data, responsesDTO.UserResponseDTO{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			PhoneNumber: user.PhoneNumber,
			Address: user.Address,
			IsActive: user.IsActive,
			CreatedAt: user.CreatedAt,
			CreatedBy: user.CreatedBy,
			UpdatedAt: user.UpdatedAt,
			UpdatedBy: user.UpdatedBy,
		})
	}
	return 200, output
}

func GetUser(GetUserRequestDTO requestsDTO.GetUserRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var user database.User
	err := db.Where("id = ?", utils.StringToUUID(GetUserRequestDTO.ID)).First(&user).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	if user.ID == (database.User{}).ID {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Data not found",
		}
		return 404, output
	}

	output := outputs.GetUserOutput{}
	output.Code = 200
	output.Message = "Success: Data found"
	output.Data = responsesDTO.UserResponseDTO{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		Address: user.Address,
		IsActive: user.IsActive,
		CreatedAt: user.CreatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedAt: user.UpdatedAt,
		UpdatedBy: user.UpdatedBy,
	}
	return 200, output
}

func RegisterUser(RegisterUserRequestDTO requestsDTO.RegisterUserRequestDTO) (int, interface{}){
	if RegisterUserRequestDTO.Password != RegisterUserRequestDTO.ConfirmPassword {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Password and Confirm Password must be same",
		}
		return 400, output
	}
	hashedPassword, err := utils.HashPassword(RegisterUserRequestDTO.Password)


	if err != nil {
		if err == context.DeadlineExceeded {
			return utils.HandleTimeout(err)
		}

		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	db := configs.GetDB()

	if db == nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error: Database connection failed",
		}
		return 500, output
	}

	user := database.User{
		Name: RegisterUserRequestDTO.Name,
		Email: RegisterUserRequestDTO.Email,
		Password: hashedPassword,
		IsActive: RegisterUserRequestDTO.IsActive,
		CreatedBy: RegisterUserRequestDTO.CreatedBy,
	}

	err = db.Create(&user).Error
	
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}
	
	output := outputs.RegisterUserOutput{}
	output.Code = 200
	output.Message = "Success: User registered"
	output.Data = responsesDTO.UserResponseDTO{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		Address: user.Address,	
		IsActive: user.IsActive,
		CreatedAt: user.CreatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedAt: user.UpdatedAt,
		UpdatedBy: user.UpdatedBy,
		}	
		return 200, output
}

func LoginUser(LoginUserRequestDTO requestsDTO.LoginUserRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var user database.User
	err := db.Where("email = ?", LoginUserRequestDTO.Email).First(&user).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	if user.ID == (database.User{}).ID {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Data not found",
		}
		return 404, output
	}

	if !utils.ComparePassword(LoginUserRequestDTO.Password, user.Password) {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Password is incorrect",
		}
		return 400, output
	}

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	output := outputs.LoginUserOutput{}
	output.Code = 200
	output.Message = "Success: User logged in"
	output.Data = responsesDTO.UserResponseDTO{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		Address: user.Address,
		IsActive: user.IsActive,
		CreatedAt: user.CreatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedAt: user.UpdatedAt,
		UpdatedBy: user.UpdatedBy,
	}
	return 200, output
}

func UpdateUser(UpdateUserRequestDTO requestsDTO.UpdateUserRequestDTO) (int, interface{}) {	
	db := configs.GetDB()
	var user database.User
	err := db.Where("id = ?", utils.StringToUUID(UpdateUserRequestDTO.ID)).First(&user).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	if user.ID == uuid.Nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Data not found",
		}
		return 404, output
	}

	if UpdateUserRequestDTO.Name != "" {
		user.Name = UpdateUserRequestDTO.Name
	}

	if UpdateUserRequestDTO.Email != "" {
		user.Email = UpdateUserRequestDTO.Email
	}

	if UpdateUserRequestDTO.PhoneNumber != "" {
		user.PhoneNumber = UpdateUserRequestDTO.PhoneNumber
	}

	if UpdateUserRequestDTO.Address != "" {
		user.Address = UpdateUserRequestDTO.Address
	}

	if err := db.Save(&user).Error; err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	output := outputs.UpdateUserOutput{}
	output.Code = 200
	output.Message = "Success: User updated"
	output.Data = responsesDTO.UserResponseDTO{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		Address: user.Address,
		IsActive: user.IsActive,
		CreatedAt: user.CreatedAt,
		CreatedBy: user.CreatedBy,
		UpdatedAt: user.UpdatedAt,
		UpdatedBy: user.UpdatedBy,
	}
	return 200, output
}
