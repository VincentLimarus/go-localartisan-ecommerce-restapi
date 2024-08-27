package helpers

import (
	"fmt"
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/outputs"
	"localArtisans/models/repositories"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"
)

func GetAllCarts(GetAllCartsRequestDTO requestsDTO.GetAllCartsRequestDTO) (int, interface{}){
	db := configs.GetDB()
	var carts []database.Carts

	if GetAllCartsRequestDTO.Limit > 100 {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit can't more than 100",
		}
		return 400, output
	}
	
	offset := (GetAllCartsRequestDTO.Page - 1) * GetAllCartsRequestDTO.Limit
	order := fmt.Sprintf("%s %s", GetAllCartsRequestDTO.OrderBy, GetAllCartsRequestDTO.OrderType)
	err := db.Offset(offset).Limit(GetAllCartsRequestDTO.Limit).Order(order).Find(&carts).Error
	
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(carts) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Carts not exist",
		}
		return 404, output
	}

	var totalData int64
	var totalPage int
	db.Model(&database.Carts{}).Count(&totalData)
	if totalData%int64(GetAllCartsRequestDTO.Limit) == 0 {
		totalPage = int(totalData / int64(GetAllCartsRequestDTO.Limit))
	} else {
		totalPage = int(totalData / int64(GetAllCartsRequestDTO.Limit)) + 1
	}

	output := outputs.GetAllCartsOutputDTO{}
	output.Page = GetAllCartsRequestDTO.Page
	output.Limit = GetAllCartsRequestDTO.Limit
	output.OrderBy = GetAllCartsRequestDTO.OrderBy
	output.OrderType = GetAllCartsRequestDTO.OrderType
	output.Code = 200
	output.Message = "Success"
	output.TotalData = int(totalData)
	output.TotalTake = len(carts)
	output.TotalPage = totalPage

	for _, cart := range carts {
		var cartInformations []responsesDTO.CartInformationResponseDTO
		cartInformations, err := repositories.GetAllCartInformationsByCartID(cart.ID.String())
		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error {Error GetAllProductByCategoryID}: " + err.Error(),
			}
			return 500, output
		}

		output.Data = append(output.Data, responsesDTO.CartResponseDTO{
			ID: cart.ID,
			UserID: cart.UserID,
			IsActive: cart.IsActive,
			CreatedBy: cart.CreatedBy,
			UpdatedBy: cart.UpdatedBy,
			CreatedAt: cart.CreatedAt,
			UpdatedAt: cart.UpdatedAt,
			CartInformations: cartInformations,
		})
	}
	return 200, output
}

func GetAllCartsByUserID(userID string) (int, interface{}){
	var carts []database.Carts
	carts, err := repositories.GetAllCartsByUserID(userID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(carts) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Carts not exist",
		}
		return 404, output
	}

	output := outputs.GetAllCartsOutputDTO{}
	output.Code = 200
	output.Message = "Success"

	for _, cart := range carts {
		var cartInformations []responsesDTO.CartInformationResponseDTO
		cartInformations, err = repositories.GetAllCartInformationsByCartID(carts[0].ID.String())
		
		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code:    500,
				Message: "Internal Server Error {Error GetAllProductByCategoryID}: " + err.Error(),
			}
			return 500, output
		}

		output.Data = append(output.Data, responsesDTO.CartResponseDTO{
			ID: cart.ID,
			UserID: cart.UserID,
			IsActive: cart.IsActive,
			CreatedBy: cart.CreatedBy,
			UpdatedBy: cart.UpdatedBy,
			CreatedAt: cart.CreatedAt,
			UpdatedAt: cart.UpdatedAt,
			CartInformations: cartInformations,
		})
	}
	return 200, output
}

func GetCartByID(cartID string) (int, interface{}){
	var cart database.Carts
	cart, err := repositories.GetCartByID(cartID)

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Category not exist",
		}
		return 404, output
	}

	var cartInformations []responsesDTO.CartInformationResponseDTO
	cartInformations, err = repositories.GetAllCartInformationsByCartID(cartID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetAllCartInformationsByCartID}: " + err.Error(),
		}
		return 500, output
	}

	output := outputs.GetCartOutputDTO{}
	output.Code = 200
	output.Message = "Success: Cart Found"
	output.Data = responsesDTO.CartResponseDTO{
		ID: cart.ID,
		UserID: cart.UserID,
		IsActive: cart.IsActive,
		CreatedBy: cart.CreatedBy,
		UpdatedBy: cart.UpdatedBy,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
		CartInformations: cartInformations,
	}
	return 200, output
}

func CreateCart(CreateCartRequestDTO requestsDTO.CreateCartRequestDTO) (int, interface{}){
	db := configs.GetDB()
	cart := database.Carts{
		UserID: CreateCartRequestDTO.UserID,
		IsActive: CreateCartRequestDTO.IsActive,
		CreatedBy: CreateCartRequestDTO.CreatedBy,
	}

	err := db.Create(&cart).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.CreateCartOutputDTO{}
	output.Code = 200
	output.Message = "Success: Cart Created"
	output.Data = responsesDTO.CartResponseDTO{
		ID: cart.ID,
		UserID: cart.UserID,
		IsActive: cart.IsActive,
		CreatedBy: cart.CreatedBy,
		UpdatedBy: cart.UpdatedBy,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}
	return 200, output
}

func DeleteCart(DeleteCartRequestDTO requestsDTO.DeleteCartRequestDTO) (int, interface{}){
	db := configs.GetDB()
	cart := database.Carts{}
	cart_information := database.CartInformations{}

	err := db.Table("carts").Where("id = ?", DeleteCartRequestDTO.ID).First(&cart).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Cart not exist",
		}
		return 404, output
	}
	
	// Delete CartInformation that linked to Cart
	err = db.Table("cart_informations").Where("cart_id = ?", DeleteCartRequestDTO.ID).Delete(&database.CartInformations{}).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	for _, cartInformation := range cart.CartInformation {
		err = db.Table("cart_information_details").Where("cart_information_id = ?", cart_information.CartID).Delete(&database.CartInformations{}).Error
		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code: 500,
				Message: "Internal Server Error" + err.Error(),
			}
			return 500, output
		}

		err := db.Delete(&cartInformation).Error

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code: 500,
				Message: "Error in Deleting References" + err.Error(),
			}
			return 500, output
		}
	}
	// End Of Deleting

	err = db.Delete(&cart).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	var cartInformations []responsesDTO.CartInformationResponseDTO
	cartInformations, err = repositories.GetAllCartInformationsByCartID(cart.ID.String())

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error {Error GetAllCartInformationsByCartID}: " + err.Error(),
		}
		return 500, output
	}
	output := outputs.DeleteCartOutputDTO{}
	output.Code = 200
	output.Message = "Success: Cart Deleted"
	output.Data = responsesDTO.CartResponseDTO{
		ID: cart.ID,
		UserID: cart.UserID,
		IsActive: cart.IsActive,
		CreatedBy: cart.CreatedBy,
		UpdatedBy: cart.UpdatedBy,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
		CartInformations: cartInformations,
	}
	return 200, output
}