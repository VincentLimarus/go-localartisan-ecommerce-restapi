package helpers

import (
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/outputs"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"
)

func GetAllCartInformations(GetAllCartInformations requestsDTO.GetAllCartInformationsRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var cartInformations []database.CartInformations	

	if GetAllCartInformations.Limit > 100 {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit can't more than 100",
		}
		return 400, output
	}
	
	offset := (GetAllCartInformations.Page - 1) * GetAllCartInformations.Limit
	order := GetAllCartInformations.OrderBy + " " + GetAllCartInformations.OrderType
	err := db.Offset(offset).Limit(GetAllCartInformations.Limit).Order(order).Find(&cartInformations).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(cartInformations) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Cart Informations not exist",
		}
		return 404, output
	}

	var totalData int64
	var totalPage int
	db.Model(&database.CartInformations{}).Count(&totalData)
	if totalData%int64(GetAllCartInformations.Limit) == 0 {
		totalPage = int(totalData / int64(GetAllCartInformations.Limit))
	} else {
		totalPage = int(totalData / int64(GetAllCartInformations.Limit)) + 1
	}

	output := outputs.GetAllCartInformationsOutput{}
	output.Page = GetAllCartInformations.Page
	output.Limit = GetAllCartInformations.Limit
	output.OrderBy = GetAllCartInformations.OrderBy	
	output.OrderType = GetAllCartInformations.OrderType
	output.Code = 200
	output.Message = "Success"
	output.TotalData = int(totalData)
	output.TotalTake = len(cartInformations)
	output.TotalPage = totalPage

	for _, cartInformation := range cartInformations {
		output.Data = append(output.Data, responsesDTO.CartInformationResponseDTO{
			ID : cartInformation.ID,
			CartID: cartInformation.CartID,
			ProductID: cartInformation.ProductID,
			Quantity: cartInformation.Quantity,
			PriceAtOrder: cartInformation.PriceAtOrder,
			IsActive: cartInformation.IsActive,
			CreatedBy: cartInformation.CreatedBy,
			UpdatedBy: cartInformation.UpdatedBy,
			CreatedAt: cartInformation.CreatedAt,
			UpdatedAt: cartInformation.UpdatedAt,
		})
	}
	return 200, output
}

func GetAllCartInformationByCartID(cartID string) (int, interface{}) {
	var cartInformations []database.CartInformations

	db := configs.GetDB()
	err := db.Table("cart_informations").Where("cart_id = ?", cartID).Find(&cartInformations).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error {Error GetAllCartInformationByCartID}: " + err.Error(),
		}
		return 500, output
	}

	if len(cartInformations) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Cart Informations not exist",
		}
		return 404, output
	}
	
	output := outputs.GetALlCartInformationByCartIDOutput{}
	output.Code = 200
	output.Message = "Success"

	for _, cartInformation := range cartInformations {
		output.Data = append(output.Data, responsesDTO.CartInformationResponseDTO{
			ID: cartInformation.ID,
			CartID: cartInformation.CartID,
			ProductID: cartInformation.ProductID,
			Quantity: cartInformation.Quantity,
			PriceAtOrder: cartInformation.PriceAtOrder,
			IsActive: cartInformation.IsActive,
			CreatedBy: cartInformation.CreatedBy,
			UpdatedBy: cartInformation.UpdatedBy,
			CreatedAt: cartInformation.CreatedAt,
			UpdatedAt: cartInformation.UpdatedAt,
		})
	}
	
	return 200, output
}

func AddItemToCart(AddItemToCart requestsDTO.AddItemToCartRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var cartInformation database.CartInformations
	var product database.Products
	// err := db.Table("products").Select("price").Where("id = ?", AddItemToCart.ProductID).First(&cartInformation.PriceAtOrder).Error 
	err := db.Table("products").Where("id = ?", AddItemToCart.ProductID).First(&product).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error disini" + err.Error(),
		}
		return 500, output
	}
	cartInformation.PriceAtOrder = product.Price

	cartInformation = database.CartInformations{
		CartID: AddItemToCart.CartID,
		ProductID: AddItemToCart.ProductID,
		Quantity: AddItemToCart.Quantity,
		PriceAtOrder: cartInformation.PriceAtOrder,
		IsActive: AddItemToCart.IsActive,
		CreatedBy: AddItemToCart.CreatedBy,
	}

	err = db.Create(&cartInformation).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.AddItemToCartOutput{}
	output.Code = 200
	output.Message = "Success: Item Added to Cart"
	output.Data = responsesDTO.CartInformationResponseDTO{
		ID : cartInformation.ID,
		CartID: AddItemToCart.CartID,
		ProductID: AddItemToCart.ProductID,
		Quantity: AddItemToCart.Quantity,
		PriceAtOrder: cartInformation.PriceAtOrder,
		IsActive: AddItemToCart.IsActive,
		CreatedBy: cartInformation.CreatedBy,
		UpdatedBy: cartInformation.UpdatedBy,
		CreatedAt: cartInformation.CreatedAt,
		UpdatedAt: cartInformation.UpdatedAt,
	}
	return 200, output
}

func UpdateItemInCart(UpdateItemInCart requestsDTO.UpdateItemInCartRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var cartInformation database.CartInformations
	
	err := db.Table("cart_informations").Where("cart_id = ? AND product_id = ?", UpdateItemInCart.CartID, UpdateItemInCart.ProductID).First(&cartInformation).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Cart Information not exist",
		}
		return 404, output
	}

	cartInformation.Quantity = UpdateItemInCart.Quantity
	
	if UpdateItemInCart.UpdatedBy == "" {
		cartInformation.UpdatedBy = "user"
	} else{
		cartInformation.UpdatedBy = UpdateItemInCart.UpdatedBy
	}

	// Boolean Update Constraint
	cartInformation.IsActive = UpdateItemInCart.IsActive

	err = db.Save(&cartInformation).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.UpdateItemInCartOutput{}
	output.Code = 200
	output.Message = "Success: Item Updated"
	output.Data = responsesDTO.CartInformationResponseDTO{
		ID : cartInformation.ID,
		CartID: cartInformation.CartID,
		ProductID: cartInformation.ProductID,
		Quantity: cartInformation.Quantity,
		PriceAtOrder: cartInformation.PriceAtOrder,
		IsActive: cartInformation.IsActive,
		CreatedBy: cartInformation.CreatedBy,
		UpdatedBy: cartInformation.UpdatedBy,
		CreatedAt: cartInformation.CreatedAt,
		UpdatedAt: cartInformation.UpdatedAt,
	}
	return 200, output
}

func DeleteItemInCart(DeleteItemInCart requestsDTO.DeleteItemInCartRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var cartInformation database.CartInformations

	err := db.Table("cart_informations").Where("cart_id = ? AND product_id = ?", DeleteItemInCart.CartID, DeleteItemInCart.ProductID).First(&cartInformation).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Cart Information not exist",
		}
		return 404, output
	}

	err = db.Delete(&cartInformation).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.DeleteItemInCartOutput{}
	output.Code = 200
	output.Message = "Success: Item Deleted"
	output.Data = responsesDTO.CartInformationResponseDTO{
		ID: cartInformation.ID,
		CartID: cartInformation.CartID,
		ProductID: cartInformation.ProductID,
		Quantity: cartInformation.Quantity,
		PriceAtOrder: cartInformation.PriceAtOrder,
		IsActive: cartInformation.IsActive,
		CreatedBy: cartInformation.CreatedBy,
		UpdatedBy: cartInformation.UpdatedBy,
		CreatedAt: cartInformation.CreatedAt,
		UpdatedAt: cartInformation.UpdatedAt,
	}

	return 200, output
}