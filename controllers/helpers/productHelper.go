package helpers

import (
	"fmt"
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/outputs"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"

	"github.com/google/uuid"
)

func GetAllProduct(GetAllProductRequestDTO requestsDTO.GetAllProductRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var products []database.Product
	
	if GetAllProductRequestDTO.Limit > 100 {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit can't more than 100",
		}
		return 400, output
	}

	offset := (GetAllProductRequestDTO.Page - 1) * GetAllProductRequestDTO.Limit
	order := fmt.Sprintf("%s %s", GetAllProductRequestDTO.OrderBy, GetAllProductRequestDTO.OrderType)
	err := db.Offset(offset).Limit(GetAllProductRequestDTO.Limit).Order(order).Find(&products).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(products) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Products not exist",
		}
		return 404, output
	}

	var totalData int64
	var totalPage int
	db.Model(&database.Product{}).Count(&totalData)
	if totalData%int64(GetAllProductRequestDTO.Limit) == 0 {
		totalPage = int(totalData / int64(GetAllProductRequestDTO.Limit))
	} else {
		totalPage = int(totalData / int64(GetAllProductRequestDTO.Limit)) + 1
	}

	output := outputs.GetAllProductOutput{}
	output.Page = GetAllProductRequestDTO.Page
	output.Limit = GetAllProductRequestDTO.Limit
	output.OrderBy = GetAllProductRequestDTO.OrderBy
	output.OrderType = GetAllProductRequestDTO.OrderType
	output.Code = 200
	output.Message = "Success: Products Found"
	output.TotalData = int(totalData)
	output.TotalPage = totalPage
	
	for _, product := range products {
		output.Data = append(output.Data, responsesDTO.ProductResponseDTO{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Quantity:    product.Quantity,
			ItemSold:    product.ItemSold,
			Rating:      product.Rating,
			CategoryID:  product.CategoryID,
			CategoryName:product.Category.Name,
			ArtisanID:   product.ArtisanID,
			ArtisanName: product.Artisans.User.Name,
			IsActive:    product.IsActive,
			CreatedBy:   product.CreatedBy,
			UpdatedBy:   product.UpdatedBy,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}
	return 200, output	
}

func GetProduct(productID string) (int, interface{}) {
	db := configs.GetDB()
	var product database.Product

	err := db.Table("products").
		Joins("JOIN artisans a ON a.id = products.artisan_id").
		Joins("JOIN categories c ON c.id = products.category_id").
		Where("products.id = ?", productID).
		Select("products.*, a.*, c.*").
		First(&product).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Product not exist",
		}
		return 404, output
	}

	output := outputs.GetProductOutput{}
	output.Code = 200
	output.Message = "Success: Product Found"
	output.Data = responsesDTO.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
		ItemSold:    product.ItemSold,
		Rating:      product.Rating,
		CategoryID:  product.CategoryID,
		CategoryName:product.Category.Name,
		ArtisanID:   product.ArtisanID,
		ArtisanName: product.Artisans.User.Name,
		IsActive:    product.IsActive,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Artisans: responsesDTO.ArtisansResponseDTO{
			ID:        product.Artisans.ID,
			UserID:    product.Artisans.UserID,
			ShopName: product.Artisans.ShopName,
			ShopAddress: product.Artisans.ShopAddress,
			Description: product.Artisans.Description,
			ShopBanner: product.Artisans.ShopBanner,
			Rating: product.Artisans.Rating,
			User: responsesDTO.UserResponseDTO{
				ID:       product.Artisans.User.ID,
				Name:     product.Artisans.User.Name,
				Email:    product.Artisans.User.Email,
				Address: product.Artisans.User.Address,
				PhoneNumber:    product.Artisans.User.PhoneNumber,
				IsActive: product.Artisans.User.IsActive,
				CreatedBy:product.Artisans.User.CreatedBy,
				UpdatedBy:product.Artisans.User.UpdatedBy,
				CreatedAt:product.Artisans.User.CreatedAt,
				UpdatedAt:product.Artisans.User.UpdatedAt,
			},
			CreatedBy: product.Artisans.CreatedBy,
			UpdatedBy: product.Artisans.UpdatedBy,	
			CreatedAt: product.Artisans.CreatedAt,
			UpdatedAt: product.Artisans.UpdatedAt,
		},
		Categories: responsesDTO.CategoryResponseDTO{
			ID:        product.Category.ID,
			Name:      product.Category.Name,
			Image:    product.Category.Image,
			IsActive:  product.Category.IsActive,
			CreatedBy:product.Category.CreatedBy,
			UpdatedBy:product.Category.UpdatedBy,
			CreatedAt:product.Category.CreatedAt,
			UpdatedAt:product.Category.UpdatedAt,
		},
	}
	return 200, output
}

func CreateProduct(CreateProductRequestDTO requestsDTO.CreateProductRequestDTO) (int, interface{}){
	db := configs.GetDB()
	product := database.Product{
		Name:        CreateProductRequestDTO.Name,
		Price:       CreateProductRequestDTO.Price,
		Description: CreateProductRequestDTO.Description,
		Quantity:    CreateProductRequestDTO.Quantity,
		CategoryID:  CreateProductRequestDTO.CategoryID,
		ArtisanID:   CreateProductRequestDTO.ArtisanID,
		CreatedBy:   CreateProductRequestDTO.CreatedBy,
		IsActive:    CreateProductRequestDTO.IsActive,
	}

	err := db.Create(&product).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.CreateProductOutput{}
	output.Code = 200
	output.Message = "Success: Product Created"
	output.Data = responsesDTO.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
		ItemSold:    product.ItemSold,
		Rating:      product.Rating,
		CategoryID:  product.CategoryID,
		CategoryName:product.Category.Name,
		ArtisanID:   product.ArtisanID,
		ArtisanName: product.Artisans.User.Name,
		IsActive:    product.IsActive,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	return 200, output
}

func UpdateProduct(UpdateProductRequestDTO requestsDTO.UpdateProductRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	product := database.Product{}

	err := db.Table("products").Where("id = ?", UpdateProductRequestDTO.ID).First(&product).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Product not exist",
		}
		return 404, output
	}

	if UpdateProductRequestDTO.Name != "" {
		product.Name = UpdateProductRequestDTO.Name
	}
	if UpdateProductRequestDTO.Price != 0 {
		product.Price = UpdateProductRequestDTO.Price
	}
	if UpdateProductRequestDTO.Description != "" {
		product.Description = UpdateProductRequestDTO.Description
	}
	if UpdateProductRequestDTO.Quantity != 0 {
		product.Quantity = UpdateProductRequestDTO.Quantity
	}
	if UpdateProductRequestDTO.CategoryID != uuid.Nil {
		product.CategoryID = UpdateProductRequestDTO.CategoryID
	}
	if UpdateProductRequestDTO.ArtisanID != uuid.Nil {
		product.ArtisanID = UpdateProductRequestDTO.ArtisanID
	}
	if UpdateProductRequestDTO.UpdatedBy == "" {
		product.UpdatedBy = "user"
	} else{
		product.UpdatedBy = UpdateProductRequestDTO.UpdatedBy
	}

	product.IsActive = UpdateProductRequestDTO.IsActive

	err = db.Save(&product).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.UpdateProductOutput{}
	output.Code = 200
	output.Message = "Success: Product Updated"
	output.Data = responsesDTO.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
		ItemSold:    product.ItemSold,
		Rating:      product.Rating,
		CategoryID:  product.CategoryID,
		CategoryName:product.Category.Name,
		ArtisanID:   product.ArtisanID,
		ArtisanName: product.Artisans.User.Name,
		IsActive:    product.IsActive,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	return 200, output
}

func DeleteProduct(DeleteProductRequestDTO requestsDTO.DeleteProductRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	product := database.Product{}

	err := db.Table("products").Where("id = ?", DeleteProductRequestDTO.ID).First(&product).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Product not exist",
		}
		return 404, output
	}

	err = db.Delete(&product).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.DeleteProductOutput{}
	output.Code = 200
	output.Message = "Success: Product Deleted"
	output.Data = responsesDTO.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
		ItemSold:    product.ItemSold,
		Rating:      product.Rating,
		CategoryID:  product.CategoryID,
		CategoryName:product.Category.Name,
		ArtisanID:   product.ArtisanID,
		ArtisanName: product.Artisans.User.Name,
		IsActive:    product.IsActive,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	return 200, output
}