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

func GetAllProduct(GetAllProductRequestDTO requestsDTO.GetAllProductRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var products []database.Products
	
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
	db.Model(&database.Products{}).Count(&totalData)
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
			ArtisanID:  product.ArtisanID,
			CategoryID: product.CategoryID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Quantity:    product.Quantity,
			ItemSold:    product.ItemSold,
			Rating:      product.Rating,
			IsActive:    product.IsActive,
			CreatedBy:   product.CreatedBy,
			UpdatedBy:   product.UpdatedBy,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}
	return 200, output	
}

func GetAllProductByArtisanID(artisanID string) (int, interface{}) {
	var products []responsesDTO.ProductResponseDTO
	products, err := repositories.GetAllProductByArtisanID(artisanID)

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Products not exist",
		}
		return 404, output
	}

	if len(products) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Products not exist",
		}
		return 404, output
	}
		
	output := outputs.GetAllProductByArtisanIDOutput{}
	output.Code = 200
	output.Message = "Success: Products Found"
	output.Data = products
	return 200, output
}

func GetAllProductByCategoryID(categoryID string) (int, interface{}) {
	var products []responsesDTO.ProductResponseDTO
	products, err := repositories.GetAllProductByCategoryID(categoryID)

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Products not exist",
		}
		return 404, output
	}

	if len(products) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Products not exist",
		}
		return 404, output
	}
	
	output := outputs.GetAllProductByCategoryIDOutput{}
	output.Code = 200
	output.Message = "Success: Products Found"
	output.Data = products
	return 200, output
}	

func GetProduct(productID string) (int, interface{}) {
	var product database.Products
	product, err := repositories.GetProductByProductID(productID)

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
		ArtisanID:  product.ArtisanID,
		CategoryID: product.CategoryID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
		ItemSold:    product.ItemSold,
		Rating:      product.Rating,
		IsActive:    product.IsActive,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	return 200, output
}

func CreateProduct(CreateProductRequestDTO requestsDTO.CreateProductRequestDTO) (int, interface{}){
	db := configs.GetDB()
	product := database.Products{
		CategoryID:  CreateProductRequestDTO.CategoryID,
		ArtisanID:   CreateProductRequestDTO.ArtisanID,
		Name:        CreateProductRequestDTO.Name,
		Price:       CreateProductRequestDTO.Price,
		Description: CreateProductRequestDTO.Description,
		Quantity:    CreateProductRequestDTO.Quantity,
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
		ArtisanID:  product.ArtisanID,
		CategoryID:  product.CategoryID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
		ItemSold:    product.ItemSold,
		Rating:      product.Rating,
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
	product := database.Products{}

	err := db.Table("products").Where("id = ?", UpdateProductRequestDTO.ID).First(&product).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Product not exist",
		}
		return 404, output
	}

	// Not NULL Update constraint -> ini tidak boleh null, kalo user tidak mengisi maka akan diisi oleh sistem
	if UpdateProductRequestDTO.Name != "" {
		product.Name = UpdateProductRequestDTO.Name
	}

	if UpdateProductRequestDTO.UpdatedBy == "" {
		product.UpdatedBy = "user"
	} else{
		product.UpdatedBy = UpdateProductRequestDTO.UpdatedBy
	}

	if UpdateProductRequestDTO.Price != 0 {
		product.Price = UpdateProductRequestDTO.Price
	}

	if UpdateProductRequestDTO.Description != "" {
		product.Description = UpdateProductRequestDTO.Description
	} 

	product.Quantity = UpdateProductRequestDTO.Quantity
	
	// Boolean Update Constraint
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
		ArtisanID:  product.ArtisanID,
		CategoryID: product.CategoryID,		
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
		ItemSold:    product.ItemSold,
		Rating:      product.Rating,
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
	product := database.Products{}

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
		ArtisanID:  product.ArtisanID,
		CategoryID: product.CategoryID,		
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Quantity:    product.Quantity,
		ItemSold:    product.ItemSold,
		Rating:      product.Rating,
		IsActive:    product.IsActive,
		CreatedBy:   product.CreatedBy,
		UpdatedBy:   product.UpdatedBy,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
	return 200, output
}

func AddProductToCart(AddProductToCartRequestDTO requestsDTO.AddProductToCartRequestDTO, LoginUser requestsDTO.LoginUserRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var user database.User
	var carts database.Carts
	var product database.Products

	err := db.Table("users").Where("email = ?", LoginUser.Email).First(&user).Error
	UserID := user.ID

	if err == nil && AddProductToCartRequestDTO.CartID == ""{
		carts = database.Carts{
			UserID: UserID,
			IsActive: AddProductToCartRequestDTO.IsActive,
			CreatedBy: AddProductToCartRequestDTO.CreatedBy,
		}	
		err = db.Create(&carts).Error

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code: 500,
				Message: "Internal Server Error" + err.Error(),
			}
			return 500, output
		}
	} 
	err = db.Table("carts").Where("user_id = ?", UserID).First(&carts).Error
	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Cart not exist",
		}
		return 404, output
	}

	err = db.Table("products").Where("id = ?", AddProductToCartRequestDTO.ID).First(&product).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Product not exist",
		}
		return 404, output
	}

	if AddProductToCartRequestDTO.Quantity > product.Quantity {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Product Quantity is not enough",
		}
		return 400, output
	}

	cartInformation := database.CartInformations{
		CartID:    carts.ID,
		ProductID:  utils.StringToUUID(AddProductToCartRequestDTO.ID),
		Quantity:  AddProductToCartRequestDTO.Quantity,
		PriceAtOrder :    product.Price,
		IsActive:  AddProductToCartRequestDTO.IsActive,
		CreatedBy: AddProductToCartRequestDTO.CreatedBy,
	}

	// if the product already exist in the cart, then the quantity will be added
	var cartInfo database.CartInformations
	err = db.Table("cart_informations").Where("cart_id = ? AND product_id = ?", carts.ID, AddProductToCartRequestDTO.ID).First(&cartInfo).Error

	if err == nil {
		cartInformation.Quantity = cartInfo.Quantity + AddProductToCartRequestDTO.Quantity
		err = db.Table("cart_informations").Where("cart_id = ? AND product_id = ?", carts.ID, AddProductToCartRequestDTO.ID).Update("quantity", cartInformation.Quantity).Error
		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code: 500,
				Message: "Internal Server Error" + err.Error(),
			}
			return 500, output
		}
		output := outputs.AddProductToCartOutput{}
		output.Code = 200
		output.Message = "Success: Product Added to Cart"

		var CartInformations []responsesDTO.CartInformationResponseDTO
		CartInformations, err = repositories.GetCartInformationByCartIDAndProductID(utils.UUIDToString(carts.ID), AddProductToCartRequestDTO.ID)

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code: 500,
				Message: "Internal Server Error" + err.Error(),
			}
			return 500, output
		}
		output.Data = responsesDTO.CartResponseDTO{
			ID:          carts.ID,
			UserID:      carts.UserID,
			IsActive:    carts.IsActive,
			CreatedBy:   carts.CreatedBy,
			CreatedAt:   carts.CreatedAt,
			UpdatedBy:   carts.UpdatedBy,
			UpdatedAt:   carts.UpdatedAt,	
			CartInformations: CartInformations,
		}
		return 200, output
	}

	err = db.Create(&cartInformation).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.AddProductToCartOutput{}
	output.Code = 200
	output.Message = "Success: Product Added to Cart"

	var CartInformations []responsesDTO.CartInformationResponseDTO
	CartInformations, err = repositories.GetCartInformationByCartIDAndProductID(utils.UUIDToString(carts.ID), AddProductToCartRequestDTO.ID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}
	output.Data = responsesDTO.CartResponseDTO{
		ID:          carts.ID,
		UserID:      carts.UserID,
		IsActive:    carts.IsActive,
		CreatedBy:   carts.CreatedBy,
		CreatedAt:   carts.CreatedAt,
		UpdatedBy:   carts.UpdatedBy,
		UpdatedAt:   carts.UpdatedAt,	
		CartInformations: CartInformations,
	}
	return 200, output
}

func CheckOutProductRequestDTO(CheckOutProductRequestDTO requestsDTO.CheckOutProductRequestDTO, LoginUser requestsDTO.LoginUserRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var user database.User
	var order database.Orders
	var product database.Products

	err := db.Table("users").Where("email = ?", LoginUser.Email).First(&user).Error
	User_ID := user.ID

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: User not exist",
		}
		return 404, output
	}

	err = db.Table("products").Where("id = ?", CheckOutProductRequestDTO.ID).First(&product).Error

	if err != nil{
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Product not exist",
		}
		return 404, output
	}

	if CheckOutProductRequestDTO.Quantity > product.Quantity {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Product Quantity is not enough",
		}
		return 400, output
	} else {
		product.Quantity = product.Quantity - CheckOutProductRequestDTO.Quantity
		err = db.Save(&product).Error

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code: 500,
				Message: "Internal Server Error" + err.Error(),
			}
			return 500, output
		}
	}
	order = database.Orders{
		UserID: User_ID,
		Status: "Waiting for Payment",
		TotalPrice: product.Price * float64(CheckOutProductRequestDTO.Quantity),
		ShippingAddress: user.Address,
		PaymentMethod: "Bank Transfer",
		IsActive: true,
	}

	err = db.Create(&order).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}
	

	orderItem := database.OrderItems{
		ProductID: product.ID,
		OrderID: order.ID,
		Quantity: CheckOutProductRequestDTO.Quantity,
		PriceAtOrder: product.Price,
		IsActive: true,
	}
	err = db.Create(&orderItem).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}	

	output := outputs.CheckoutProductOutput{}
	output.Code = 200
	output.Message = "Success: Product Checked Out"

	var orderItems []responsesDTO.OrderItemsResponseDTO
	orderItems, err = repositories.GetAllOrderItemsByOrderID(order.ID.String())

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	output.Data = responsesDTO.OrderResponseDTO{
		ID:              order.ID,
		UserID:          order.UserID,
		Status:          order.Status,
		TotalPrice:      order.TotalPrice,
		ShippingAddress: order.ShippingAddress,
		PaymentMethod:   order.PaymentMethod,
		IsActive:        order.IsActive,
		CreatedBy:       order.CreatedBy,
		UpdatedBy:       order.UpdatedBy,
		CreatedAt:       order.CreatedAt,
		UpdatedAt:       order.UpdatedAt,
		OrderItems: 	 orderItems,
	}
	return 200, output
}