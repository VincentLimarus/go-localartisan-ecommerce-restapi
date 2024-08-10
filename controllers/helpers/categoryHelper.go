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
)

func GetAllCategories(GetAllCategoriesRequestDTO requestsDTO.GetAllCategoriesRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var categories []database.Categories
	
	if GetAllCategoriesRequestDTO.Limit > 100 {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit can't more than 100",
		}
		return 400, output
	}

	offset := (GetAllCategoriesRequestDTO.Page - 1) * GetAllCategoriesRequestDTO.Limit
	order := fmt.Sprintf("%s %s", GetAllCategoriesRequestDTO.OrderBy, GetAllCategoriesRequestDTO.OrderType)
	err := db.Offset(offset).Limit(GetAllCategoriesRequestDTO.Limit).Order(order).Find(&categories).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(categories) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Categories not exist",
		}
		return 404, output
	}

	var totalData int64
	var totalPage int
	db.Model(&database.Categories{}).Count(&totalData)
	if totalData%int64(GetAllCategoriesRequestDTO.Limit) == 0 {
		totalPage = int(totalData / int64(GetAllCategoriesRequestDTO.Limit))
	} else {
		totalPage = int(totalData / int64(GetAllCategoriesRequestDTO.Limit)) + 1
	}

	output := outputs.GetAllCategoryOutput{}
	output.Page = GetAllCategoriesRequestDTO.Page
	output.Limit = GetAllCategoriesRequestDTO.Limit
	output.OrderBy = GetAllCategoriesRequestDTO.OrderBy
	output.OrderType = GetAllCategoriesRequestDTO.OrderType
	output.Code = 200
	output.Message = "Success: Categories Found"
	output.TotalData = int(totalData)
	output.TotalTake = len(categories)
	output.TotalPage = totalPage

	for _, category := range categories {
		var productResponseDTO []responsesDTO.ProductResponseDTO
		productResponseDTO, err = repositories.GetAllProductByCategoryID(category.ID.String())

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error {Error GetAllProductByArtisanID}: " + err.Error(),
			}
			return 500, output
		}
		output.Data = append(output.Data, responsesDTO.CategoryResponseDTO{
			ID:        category.ID,
			Name:      category.Name,
			Image:     category.Image,
			IsActive:  category.IsActive,
			CreatedBy: category.CreatedBy,
			UpdatedBy: category.UpdatedBy,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
			Products: productResponseDTO,
		})
	}
	return 200, output
}

func GetCategory(categoryID string) (int, interface{}){
	var category database.Categories

	category, err := repositories.GetCategoryByCategoryID(categoryID)

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Category not exist",
		}
		return 404, output
	}

	var productResponseDTO []responsesDTO.ProductResponseDTO
	productResponseDTO, err = repositories.GetAllProductByCategoryID(categoryID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetAllProductByArtisanID}: " + err.Error(),
		}
		return 500, output
	}
	
	output := outputs.GetCategoryOutput{}
	output.Code = 200
	output.Message = "Success: Category Found"
	output.Data = responsesDTO.CategoryResponseDTO{
		ID:        category.ID,
		Name:      category.Name,
		Image:     category.Image,
		IsActive:  category.IsActive,
		CreatedBy: category.CreatedBy,
		UpdatedBy: category.UpdatedBy,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Products: productResponseDTO,
	}
	return 200, output
}

func CreateCategory(CreateCategoryRequestDTO requestsDTO.CreateCategoryRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	category := database.Categories{
		Name:      CreateCategoryRequestDTO.Name,
		Image:     CreateCategoryRequestDTO.Image,
		IsActive:  CreateCategoryRequestDTO.IsActive,
		CreatedBy: CreateCategoryRequestDTO.CreatedBy,
	}

	err := db.Create(&category).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.CreateCategoryOutput{}
	output.Code = 200
	output.Message = "Success: Category Created"
	output.Data = responsesDTO.CategoryResponseDTO{
		ID:        category.ID,
		Name:      category.Name,
		Image:     category.Image,
		IsActive:  category.IsActive,
		CreatedBy: category.CreatedBy,
		UpdatedBy: category.UpdatedBy,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
	return 200, output
}

func UpdateCategory(UpdateCategoryRequestDTO requestsDTO.UpdateCategoryRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	category := database.Categories{}
	
	err := db.Table("categories").Where("id = ?", UpdateCategoryRequestDTO.ID).First(&category).Error
	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Category not exist",
		}
		return 404, output
	}

	if UpdateCategoryRequestDTO.Name != "" {
		category.Name = UpdateCategoryRequestDTO.Name
	}
	if UpdateCategoryRequestDTO.Image != "" {
		category.Image = UpdateCategoryRequestDTO.Image
	}
	if UpdateCategoryRequestDTO.UpdatedBy == "" {
		category.UpdatedBy = "user"
	} else{
		category.UpdatedBy = UpdateCategoryRequestDTO.UpdatedBy
	}

	category.IsActive = UpdateCategoryRequestDTO.IsActive

	err = db.Save(&category).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	var productResponseDTO []responsesDTO.ProductResponseDTO
	productResponseDTO, err = repositories.GetAllProductByCategoryID(UpdateCategoryRequestDTO.ID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetAllProductByArtisanID}: " + err.Error(),
		}
		return 500, output
	}

	output := outputs.UpdateCategoryOutput{}
	output.Code = 200
	output.Message = "Success: Category Updated"
	output.Data = responsesDTO.CategoryResponseDTO{
		ID:        category.ID,
		Name:      category.Name,
		Image:     category.Image,
		IsActive:  category.IsActive,
		CreatedBy: category.CreatedBy,
		UpdatedBy: category.UpdatedBy,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Products: productResponseDTO,
	}
	return 200, output
}

func DeleteCategory(DeleteCategoryRequestDTO requestsDTO.DeleteCategoryRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var category database.Categories
	var products []database.Products
	
	err := db.Where("category_id = ?", utils.StringToUUID(DeleteCategoryRequestDTO.ID)).Find(&products).Error

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


	err = db.Table("categories").Where("id = ?", DeleteCategoryRequestDTO.ID).First(&category).Error
	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Category not exist",
		}
		return 404, output
	}

	err = db.Delete(&category).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	var productResponseDTO []responsesDTO.ProductResponseDTO
	productResponseDTO, err = repositories.GetAllProductByCategoryID(DeleteCategoryRequestDTO.ID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetAllProductByArtisanID}: " + err.Error(),
		}
		return 500, output
	}

	output := outputs.DeleteCategoryOutput{}
	output.Code = 200
	output.Message = "Success: Category Deleted"
	output.Data = responsesDTO.CategoryResponseDTO{
		ID:        category.ID,
		Name:      category.Name,
		Image:     category.Image,
		IsActive:  category.IsActive,
		CreatedBy: category.CreatedBy,
		UpdatedBy: category.UpdatedBy,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
		Products: productResponseDTO,
	}
	return 200, output
}