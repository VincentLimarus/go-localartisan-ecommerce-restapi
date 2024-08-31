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

func GetAllCarts(c *gin.Context){
	var GetAllCartsRequestDTO requestsDTO.GetAllCartsRequestDTO
	GetAllCartsRequestDTO.Page, GetAllCartsRequestDTO.Limit, GetAllCartsRequestDTO.OrderBy, GetAllCartsRequestDTO.OrderType = utils.PaginationHandler(GetAllCartsRequestDTO.Page, GetAllCartsRequestDTO.Limit, GetAllCartsRequestDTO.OrderBy, GetAllCartsRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllCartsRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllCarts(GetAllCartsRequestDTO)
	c.JSON(code, output)
}

func GetAllCartByUserID(c *gin.Context){
	cartID := c.Param("id")

	if _, err := uuid.Parse(cartID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllCartsByUserID(cartID)
	c.JSON(code, output)
}

func GetCartByID(c *gin.Context){
	cartID := c.Param("id")

	if _, err := uuid.Parse(cartID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetCartByID(cartID)
	c.JSON(code, output)
}

func CreateCart(c *gin.Context){
	var CreateCartRequestDTO requestsDTO.CreateCartRequestDTO
	if err := c.ShouldBindJSON(&CreateCartRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.CreateCart(CreateCartRequestDTO)
	c.JSON(code, output)
}

func DeleteCart(c *gin.Context){
	var DeleteCartRequestDTO requestsDTO.DeleteCartRequestDTO
	if err := c.ShouldBindJSON(&DeleteCartRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteCart(DeleteCartRequestDTO)
	c.JSON(code, output)
}

func CheckoutProductFromCart(c *gin.Context) {
	var CheckoutProductFromCart requestsDTO.CheckoutProductFromCartRequestDTO

	// Bind JSON body ke CheckoutProductRequestDTO
	if err := c.ShouldBindJSON(&CheckoutProductFromCart); err != nil {
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
	code, output := helpers.CheckoutProductFromCart(CheckoutProductFromCart, LoginUser)
	c.JSON(code, output)

}

func AuthCartService(router *gin.RouterGroup) {
	router.GET("/carts", GetAllCarts)
	router.GET("/carts/user/:id", GetAllCartByUserID)
	router.GET("/cart/:id", GetCartByID)
	router.POST("/cart/create", CreateCart)
	router.POST("/cart/delete", DeleteCart)
	router.POST("/cart/checkout", CheckoutProductFromCart)
}