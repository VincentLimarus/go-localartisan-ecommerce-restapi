package services

import (
	"fmt"
	"localArtisans/controllers/helpers"
	"localArtisans/models/outputs"
	"localArtisans/models/requestsDTO"
	"localArtisans/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

func GetAllProduct(c *gin.Context){
	var GetAllProductRequestDTO requestsDTO.GetAllProductRequestDTO
	GetAllProductRequestDTO.Page, GetAllProductRequestDTO.Limit, GetAllProductRequestDTO.OrderBy, GetAllProductRequestDTO.OrderType = utils.PaginationHandler(GetAllProductRequestDTO.Page, GetAllProductRequestDTO.Limit, GetAllProductRequestDTO.OrderBy, GetAllProductRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllProductRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllProduct(GetAllProductRequestDTO)
	c.JSON(code, output)
}

func GetProductByID(c *gin.Context){
	productID := c.Param("id")

	if _, err := uuid.Parse(productID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetProduct(productID)
	c.JSON(code, output)
}

func CreateProduct(c *gin.Context){
	var CreateProductRequestDTO requestsDTO.CreateProductRequestDTO

	if err := c.ShouldBindJSON(&CreateProductRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.CreateProduct(CreateProductRequestDTO)
	c.JSON(code, output)
}

func UpdateProduct(c *gin.Context){
	var UpdateProductRequestDTO requestsDTO.UpdateProductRequestDTO

	if err := c.ShouldBindJSON(&UpdateProductRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.UpdateProduct(UpdateProductRequestDTO)
	c.JSON(code, output)
}

func DeleteProduct(c *gin.Context){
	var DeleteProductRequestDTO requestsDTO.DeleteProductRequestDTO

	if err := c.ShouldBindJSON(&DeleteProductRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteProduct(DeleteProductRequestDTO)
	c.JSON(code, output)
}

func GetAllProductByArtisanID(c *gin.Context){
	artisanID := c.Param("id")

	if _, err := uuid.Parse(artisanID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllProductByArtisanID(artisanID)
	c.JSON(code, output)
}

func GetAllProductByCategoryID(c *gin.Context){	
	categoryID := c.Param("id")

	if _, err := uuid.Parse(categoryID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllProductByCategoryID(categoryID)
	c.JSON(code, output)
}

func AddProductToCart(c *gin.Context) {
	var AddProductToCart requestsDTO.AddProductToCartRequestDTO

	// Bind JSON body ke AddProductToCartRequestDTO
	if err := c.ShouldBindJSON(&AddProductToCart); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}

	// Ambil email dari context
	email, exists := c.Get("user_email")
	if !exists {
		output := outputs.UnauthorizedOutput{
			Code:    401,
			Message: "Unauthorized: User email not found in context",
		}
		c.JSON(http.StatusUnauthorized, output)
		return
	}

	// Buat LoginUserRequestDTO dengan email yang diambil dari context
	LoginUser := requestsDTO.LoginUserRequestDTO{
		Email: email.(string),
	}

	// Panggil fungsi helper dengan data yang telah di-bind dan LoginUser
	code, output := helpers.AddProductToCart(AddProductToCart, LoginUser)
	c.JSON(code, output)
}

func BaseProductService(router *gin.RouterGroup) {
	router.GET("/products", GetAllProduct)
	router.GET("/products/artisan/:id", GetAllProductByArtisanID)
	router.GET("/products/category/:id", GetAllProductByCategoryID)
	router.GET("/product/:id", GetProductByID)
}

func AuthProductService(router *gin.RouterGroup) {
	router.POST("/product/create", CreateProduct)
	router.POST("/product/update", UpdateProduct)
	router.POST("/product/delete", DeleteProduct)
	router.POST("/product/add-to-cart", AddProductToCart)
}
