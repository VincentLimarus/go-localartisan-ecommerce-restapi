package helpers

import (
	"fmt"
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/outputs"
	"localArtisans/models/repositories"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"
	"localArtisans/utils"

	"github.com/google/uuid"
)

func GetAllArtisans(GetAllArtisansRequestDTO requestsDTO.GetAllArtisansRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var artisans []database.Artisans

	if GetAllArtisansRequestDTO.Limit > 100{
		output := outputs.BadRequestOutput{
			Code : 400,
			Message: "Bad Reuqest: Limit can't more than 100",
		}
		return 400, output
	}

	offset := (GetAllArtisansRequestDTO.Page - 1) * GetAllArtisansRequestDTO.Limit
	order := fmt.Sprintf("%s %s", GetAllArtisansRequestDTO.OrderBy, GetAllArtisansRequestDTO.OrderType)
	err := db.Offset(offset).Limit(GetAllArtisansRequestDTO.Limit).Order(order).Find(&artisans).Error


	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(artisans) == 0{
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Artisans not exist",
		}
		return 404, output
	}
	var totalData int64
	var totalPage int 

	db.Model(&database.Artisans{}).Count(&totalData)

	if totalData % int64(GetAllArtisansRequestDTO.Limit) == 0 {
		totalPage = int(totalData / int64(GetAllArtisansRequestDTO.Limit))
	} else{
		totalPage = int(totalData / int64(GetAllArtisansRequestDTO.Limit)) + 1
	}

	output := outputs.GetAllArtisanOutput{}
	output.Page = GetAllArtisansRequestDTO.Page
	output.Limit = GetAllArtisansRequestDTO.Limit
	output.OrderBy = GetAllArtisansRequestDTO.OrderBy
	output.OrderType = GetAllArtisansRequestDTO.OrderType
	output.Code = 200
	output.Message = "Success: Artisans Found"
	output.TotalData = int(totalData)
	output.TotalTake = len(artisans)
	output.TotalPage = totalPage

	for _, artisan := range artisans{
		output.Data = append(output.Data, responsesDTO.ArtisansResponseDTO{
			ID: artisan.ID,
			UserID: artisan.UserID,
			ShopName: artisan.ShopName,
			ShopAddress: artisan.ShopAddress,
			Description: artisan.Description,
			ShopBanner: artisan.ShopBanner,
			Rating: artisan.Rating,
			IsActive: artisan.IsActive,
			CreatedBy: artisan.CreatedBy,
			UpdatedBy: artisan.UpdatedBy,
			CreatedAt: artisan.CreatedAt,
			UpdatedAt: artisan.UpdatedAt,
		})
	}
	return 200, output
}

func GetArtisan(artisanID string) (int, interface{}) {
	// var user database.User
	var artisan database.Artisans

	artisan, err := repositories.GetArtisanByArtisanID(artisanID)
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetArtisanByArtisanID}: " + err.Error(),
		}
		return 500, output
	}

	if artisan.ID == uuid.Nil {
		output := outputs.NotFoundOutput{
			Code:    404,
			Message: "Not Found: Artisan does not exist",
		}
		return 404, output
	}

	// user, err = repositories.GetUserByUserID(artisanID)
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetUserByUserID}: " + err.Error(),
		}
		return 500, output
	}

	// joinErr := db.Table("artisans").
	// 	Joins("JOIN users ON users.id = artisans.user_id").
	// 	Where("artisans.id = ?", utils.StringToUUID(artisanID)).
	// 	Select("artisans.*, users.id as user_id, users.name as user_name, users.email as user_email, users.phone_number as user_phone, users.address as user_address, users.is_active as user_is_active, users.created_by as user_created_by, users.updated_by as user_updated_by, users.created_at as user_created_at, users.updated_at as user_updated_at").
	// 	First(&artisan).Error
	
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error: " + err.Error(),
		}
		return 500, output
	}

	output := outputs.GetArtisanOutput{}
	output.Code = 200
	output.Message = "Success: Artisan Found"
	output.Data = responsesDTO.ArtisansResponseDTO{
		ID:          artisan.ID,
		UserID:      artisan.UserID,
		ShopName:    artisan.ShopName,
		ShopAddress: artisan.ShopAddress,
		Description: artisan.Description,
		ShopBanner:  artisan.ShopBanner,
		Rating:      artisan.Rating,
		IsActive:    artisan.IsActive,
		CreatedBy:   artisan.CreatedBy,
		UpdatedBy:   artisan.UpdatedBy,
		CreatedAt:   artisan.CreatedAt,
		UpdatedAt:   artisan.UpdatedAt,

	}
	return 200, output
}

func RegisterArtisan(RegisterArtisanRequestDTO requestsDTO.RegisterArtisanRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	
	artisan := database.Artisans{
		UserID: RegisterArtisanRequestDTO.UserID,
		ShopName: RegisterArtisanRequestDTO.ShopName,
		ShopAddress: RegisterArtisanRequestDTO.ShopAddress,
		Description: RegisterArtisanRequestDTO.Description,
		ShopBanner: RegisterArtisanRequestDTO.ShopBanner,
		IsActive: RegisterArtisanRequestDTO.IsActive,
		CreatedBy: RegisterArtisanRequestDTO.CreatedBy,
		UpdatedBy: RegisterArtisanRequestDTO.CreatedBy,
	}

	err := db.Create(&artisan).Error

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.RegisterArtisanOutput{}
	output.Code = 200
	output.Message = "Success: Artisan Registered"
	output.Data = responsesDTO.ArtisansResponseDTO{
		ID: artisan.ID,
		UserID: artisan.UserID,
		ShopName: artisan.ShopName,
		ShopAddress: artisan.ShopAddress,
		Description: artisan.Description,
		ShopBanner: artisan.ShopBanner,
		Rating: artisan.Rating,
		IsActive: artisan.IsActive,
		CreatedBy: artisan.CreatedBy,
		UpdatedBy: artisan.UpdatedBy,
		CreatedAt: artisan.CreatedAt,
		UpdatedAt: artisan.UpdatedAt,

	}
	return 200, output
}

func UpdateArtisan(UpdateArtisanRequestDTO requestsDTO.UpdateArtisanRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var artisan database.Artisans
	var user database.User

	joinErr := db.Table("users").
		Joins("JOIN artisans ON users.id = artisans.user_id").
		Where("users.id = ?", UpdateArtisanRequestDTO.UserID).
		First(&user).Error

	if joinErr != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error {Error Join Query}" + joinErr.Error(),
		}
		return 500, output
	}
	
	err := db.Where("id = ?", utils.StringToUUID(UpdateArtisanRequestDTO.ID)).First(&artisan).Error
	
	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if artisan.ID == (database.Artisans{}).ID{
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Artisan not exist",
		}
		return 404, output
	}

	artisan.IsActive = UpdateArtisanRequestDTO.IsActive
	
	if UpdateArtisanRequestDTO.UserID != uuid.Nil{
		artisan.UserID = UpdateArtisanRequestDTO.UserID
	}
	if UpdateArtisanRequestDTO.ShopName != ""{
		artisan.ShopName = UpdateArtisanRequestDTO.ShopName
	}
	if UpdateArtisanRequestDTO.ShopAddress != ""{
		artisan.ShopAddress = UpdateArtisanRequestDTO.ShopAddress
	}
	if UpdateArtisanRequestDTO.Description != ""{
		artisan.Description = UpdateArtisanRequestDTO.Description
	}
	if UpdateArtisanRequestDTO.ShopBanner != ""{
		artisan.ShopBanner = UpdateArtisanRequestDTO.ShopBanner
	}
	if UpdateArtisanRequestDTO.UpdatedBy == ""{
		artisan.UpdatedBy = "user"
	}else{
		artisan.UpdatedBy = UpdateArtisanRequestDTO.UpdatedBy
	}
	
	err = db.Save(&artisan).Error
	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.UpdateArtisanOutput{}
	output.Code = 200
	output.Message = "Success: Artisan Updated"
	output.Data = responsesDTO.ArtisansResponseDTO{
		ID: artisan.ID,
		UserID: artisan.UserID,
		ShopName: artisan.ShopName,
		ShopAddress: artisan.ShopAddress,
		Description: artisan.Description,
		ShopBanner: artisan.ShopBanner,
		Rating: artisan.Rating,
		IsActive: artisan.IsActive,
		CreatedBy: artisan.CreatedBy,
		UpdatedBy: artisan.UpdatedBy,
		CreatedAt: artisan.CreatedAt,
		UpdatedAt: artisan.UpdatedAt,

	}
	return 200, output
}

func DeleteArtisan(DeleteArtisanRequestDTO requestsDTO.DeleteArtisanRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var artisan database.Artisans
	var user database.User

	joinErr := db.Table("users").
		Joins("JOIN artisans ON users.id = artisans.user_id").
		Where("users.id = ?", DeleteArtisanRequestDTO.UserID).
		First(&user).Error

	if joinErr != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error {Error Join Query} " + joinErr.Error(),
		}
		return 500, output
	}	

	err := db.Where("id = ?", utils.StringToUUID(DeleteArtisanRequestDTO.ID)).First(&artisan).Error

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if artisan.ID == (database.Artisans{}).ID{
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Artisan not exist",
		}
		return 404, output
	}

	err = db.Delete(&artisan).Error
	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.DeleteArtisanOutput{}
	output.Code = 200
	output.Message = "Success: Artisan Deleted"
	output.Data = responsesDTO.ArtisansResponseDTO{
		ID: artisan.ID,
		UserID: artisan.UserID,
		ShopName: artisan.ShopName,
		ShopAddress: artisan.ShopAddress,
		Description: artisan.Description,
		ShopBanner: artisan.ShopBanner,
		Rating: artisan.Rating,
		IsActive: artisan.IsActive,
		CreatedBy: artisan.CreatedBy,
		UpdatedBy: artisan.UpdatedBy,
		CreatedAt: artisan.CreatedAt,
		UpdatedAt: artisan.UpdatedAt,

	}
	
	return 200, output
}