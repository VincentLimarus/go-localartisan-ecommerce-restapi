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

func ProductService(router *gin.RouterGroup) {
	router.GET("/products", GetAllProduct)
	router.GET("/product/:id", GetProductByID)
	router.POST("/product/create", CreateProduct)
	router.POST("/product/update", UpdateProduct)
	router.POST("/product/delete", DeleteProduct)
}