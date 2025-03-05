package helpers

import (
	"context"
	"fmt"
	"localArtisans/configs"
	"localArtisans/models/database"
	outputs "localArtisans/models/outputs"
	"localArtisans/models/repositories"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"
	"localArtisans/utils"

	"github.com/google/uuid"
)

func GetAllUser(GetAllUsersRequestDTO requestsDTO.GetAllUsersRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var users []database.User

	if GetAllUsersRequestDTO.Limit == 0  || GetAllUsersRequestDTO.Limit > 100{
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit cannot be 0 / 100",
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
		userResponse := responsesDTO.UserResponseDTO{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			PhoneNumber: user.PhoneNumber,
			Address: user.Address,
			IsArtisan: user.IsArtisan,
			IsActive: user.IsActive,
			CreatedBy: user.CreatedBy,
			UpdatedBy: user.UpdatedBy,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		output.Data = append(output.Data, userResponse)
	}
	
	return 200, output
}


func GetUser(userID string) (int, interface{}) {
	var user database.User
	user, err := repositories.GetUserByUserID(userID)

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

	output := outputs.GetUserOutput{
		BaseOutput: outputs.BaseOutput{
			Code: 200,
			Message: "Success: Data found",
		},
		Data: responsesDTO.UserResponseDTO{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			PhoneNumber: user.PhoneNumber,
			Address: user.Address,
			IsArtisan: user.IsArtisan,
			IsActive: user.IsActive,
			CreatedBy: user.CreatedBy,
			UpdatedBy: user.UpdatedBy,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
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

	user := database.User{
		Name: RegisterUserRequestDTO.Name,
		Email: RegisterUserRequestDTO.Email,
		Password: hashedPassword,
		IsArtisan: false,
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
		IsArtisan: user.IsArtisan,
		IsActive: user.IsActive,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		}	
		return 200, output
}

func LoginUser(LoginUserRequestDTO requestsDTO.LoginUserRequestDTO) (int, interface{}, string) {
	var user database.User
	user, err := repositories.GetUserByEmail(LoginUserRequestDTO.Email)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output, err.Error()
	}

	if !utils.ComparePassword(user.Password, LoginUserRequestDTO.Password) {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: Password or Email is incorrect",
		}
		return 400, output, "Token not found"
	}

	token, tokenErr := utils.CreateJWTToken(user.ID, user.Email)

	if tokenErr != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: fmt.Sprintf("Internal Server Error: %v", tokenErr),
		}
		return 500, output, tokenErr.Error()
	}

	output := outputs.LoginUserOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: User logged in",
		},
		Data: responsesDTO.UserResponseDTO{
			ID:           user.ID,
			Name:         user.Name,
			Email:        user.Email,
			PhoneNumber:  user.PhoneNumber,
			Address:      user.Address,
			IsArtisan:    user.IsArtisan,
			IsActive:     user.IsActive,
			CreatedBy:    user.CreatedBy,
			UpdatedBy:    user.UpdatedBy,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		},
	}

	return 200, output, token
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

	// Not NULL Update constraint -> ini tidak boleh null, kalo user tidak mengisi maka akan diisi oleh sistem
	if UpdateUserRequestDTO.Name != "" {
		user.Name = UpdateUserRequestDTO.Name
	}

	if UpdateUserRequestDTO.Email != "" {
		user.Email = UpdateUserRequestDTO.Email
	} 
	
	if UpdateUserRequestDTO.UpdatedBy == "" {
		user.UpdatedBy = "User"
	} 

	// Nullable Update Constraint
	user.PhoneNumber = UpdateUserRequestDTO.PhoneNumber	
	user.Address = UpdateUserRequestDTO.Address	

	// Boolean Update Constraint
	user.IsActive = UpdateUserRequestDTO.IsActive

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
		IsArtisan: user.IsArtisan,
		IsActive: user.IsActive,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return 200, output
}

func DeleteUser(DeleteUserRequestDTO requestsDTO.DeleteUserRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var user database.User
	
	err := db.Where("id = ?", utils.StringToUUID(DeleteUserRequestDTO.ID)).First(&user).Error

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

	if !utils.ComparePassword(user.Password, DeleteUserRequestDTO.Password) {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Password is incorrect",
		}
		return 400, output
	}

	output := outputs.DeleteUserOutput{}
	output.Code = 200
	output.Message = "Success: User deleted"
	output.Data = responsesDTO.UserResponseDTO{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		Address: user.Address,
		IsArtisan: user.IsArtisan,
		IsActive: user.IsActive,
		CreatedBy: user.CreatedBy,
		UpdatedBy: user.UpdatedBy,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	
	if err := db.Delete(&user).Error; err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	return 200, output
}

func ChangePasswordUser(ChangePasswordUserRequestDTO requestsDTO.ChangePasswordRequestDTO) (int, interface{}){
	db := configs.GetDB()
	var user database.User
	err := db.Where("id = ?", utils.StringToUUID(ChangePasswordUserRequestDTO.ID)).First(&user).Error

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

	if !utils.ComparePassword(user.Password, ChangePasswordUserRequestDTO.OldPassword) {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Old Password is incorrect",
		}
		return 400, output
	}

	if ChangePasswordUserRequestDTO.NewPassword != ChangePasswordUserRequestDTO.ConfirmPassword {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: New Password and Confirm Password must be same",
		}
		return 400, output
	}

	hashedPassword, err := utils.HashPassword(ChangePasswordUserRequestDTO.NewPassword)

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

	user.Password = hashedPassword
	user.UpdatedBy = "User" // Updated being omitempty in the struct so if the user updated the password, user won't fill the JSON so it will automatically filled by the system as (user)

	if err := db.Save(&user).Error; err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		}
		return 500, output
	}

	output := outputs.ChangePasswordOutput{}
	output.Code = 200
	output.Message = "Success: Password changed"

	return 200, output
}