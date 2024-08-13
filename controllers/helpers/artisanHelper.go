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
		var productResponseDTO []responsesDTO.ProductResponseDTO
		productResponseDTO, err = repositories.GetAllProductByArtisanID(artisan.ID.String())

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error {Error GetAllProductByArtisanID}: " + err.Error(),
			}
			return 500, output
		}
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
			Products: productResponseDTO,
		})
	}
	return 200, output
}

func GetArtisan(artisanID string) (int, interface{}) {
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

	var productResponseDTO []responsesDTO.ProductResponseDTO
	productResponseDTO, err = repositories.GetAllProductByArtisanID(artisanID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetAllProductByArtisanID}: " + err.Error(),
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
		Products:   productResponseDTO,

	}
	return 200, output
}

func RegisterArtisan(RegisterArtisanRequestDTO requestsDTO.RegisterArtisanRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var artisan database.Artisans

	err := db.Table("artisans").Where("user_id = ?", RegisterArtisanRequestDTO.UserID).First(&artisan).Error

	if err == nil{
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: User already registered as artisan (1 User can only be 1 Artisan)",
		}
		return 400, output
	}

	artisan = database.Artisans{
		UserID: RegisterArtisanRequestDTO.UserID,
		ShopName: RegisterArtisanRequestDTO.ShopName,
		ShopAddress: RegisterArtisanRequestDTO.ShopAddress,
		Description: RegisterArtisanRequestDTO.Description,
		ShopBanner: RegisterArtisanRequestDTO.ShopBanner,
		IsActive: RegisterArtisanRequestDTO.IsActive,
		CreatedBy: RegisterArtisanRequestDTO.CreatedBy,
		UpdatedBy: RegisterArtisanRequestDTO.CreatedBy,
	}

	err = db.Create(&artisan).Error

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	var user database.User
	user, err = repositories.GetUserByUserID(artisan.UserID.String()) 

	if user.ID == RegisterArtisanRequestDTO.UserID && err == nil{
		user.IsArtisan = true
		err = db.Save(&user).Error
		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error {Error Save User}: " + err.Error(),
			}
			return 500, output
		}
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

	// Not NULL Update constraint -> ini tidak boleh null, kalo user tidak mengisi maka akan diisi oleh sistem
	if UpdateArtisanRequestDTO.ShopName != ""{
		artisan.ShopName = UpdateArtisanRequestDTO.ShopName
	}
	if UpdateArtisanRequestDTO.ShopAddress != ""{
		artisan.ShopAddress = UpdateArtisanRequestDTO.ShopAddress
	}
	if UpdateArtisanRequestDTO.UpdatedBy == ""{
		artisan.UpdatedBy = "user"
	}else{
		artisan.UpdatedBy = UpdateArtisanRequestDTO.UpdatedBy
	}

	artisan.Description = UpdateArtisanRequestDTO.Description
	artisan.ShopBanner = UpdateArtisanRequestDTO.ShopBanner
	artisan.IsActive = UpdateArtisanRequestDTO.IsActive
	
	err = db.Save(&artisan).Error
	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	var productResponseDTO []responsesDTO.ProductResponseDTO
	productResponseDTO, err = repositories.GetAllProductByArtisanID(UpdateArtisanRequestDTO.ID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetAllProductByArtisanID}: " + err.Error(),
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
		Products: productResponseDTO,
	}
	return 200, output
}

func DeleteArtisan(DeleteArtisanRequestDTO requestsDTO.DeleteArtisanRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var artisan database.Artisans
	var products []database.Products

	err := db.Where("artisan_id = ?", utils.StringToUUID(DeleteArtisanRequestDTO.ID)).Find(&products).Error

	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	for _, product := range products{
		if err := db.Delete(&product).Error; err != nil {
			output := outputs.InternalServerErrorOutput{
				Code: 500,
				Message: fmt.Sprintf("Internal Server Error: %v", err),
			}
			return 500, output
		}
	}

	err = db.Where("id = ?", utils.StringToUUID(DeleteArtisanRequestDTO.ID)).First(&artisan).Error

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

	var user database.User
	user, err = repositories.GetUserByUserID(DeleteArtisanRequestDTO.UserID.String())

	if user.ID == DeleteArtisanRequestDTO.UserID && err == nil{
		user.IsArtisan = false
		err = db.Save(&user).Error
		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error {Error Save User}: " + err.Error(),
			}
			return 500, output
		}
	}

	err = db.Delete(&artisan).Error
	if err != nil{
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	var productResponseDTO []responsesDTO.ProductResponseDTO
	productResponseDTO, err = repositories.GetAllProductByArtisanID(DeleteArtisanRequestDTO.ID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetAllProductByArtisanID}: " + err.Error(),
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
		Products: productResponseDTO,
	}
	return 200, output
}