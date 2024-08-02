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

func GetAllCategories(c *gin.Context){
	var GetAllCategoriesRequestDTO requestsDTO.GetAllCategoriesRequestDTO
	GetAllCategoriesRequestDTO.Page, GetAllCategoriesRequestDTO.Limit, GetAllCategoriesRequestDTO.OrderBy, GetAllCategoriesRequestDTO.OrderType = utils.PaginationHandler(GetAllCategoriesRequestDTO.Page, GetAllCategoriesRequestDTO.Limit, GetAllCategoriesRequestDTO.OrderBy, GetAllCategoriesRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllCategoriesRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllCategories(GetAllCategoriesRequestDTO)
	c.JSON(code, output)
}

func GetCategory(c *gin.Context){
	categoryID := c.Param("id")

	if _, err := uuid.Parse(categoryID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetCategory(categoryID)
	c.JSON(code, output)
}

func CreateCategory(c *gin.Context){
	var CreateCategoryRequestDTO requestsDTO.CreateCategoryRequestDTO

	if err := c.ShouldBindJSON(&CreateCategoryRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.CreateCategory(CreateCategoryRequestDTO)
	c.JSON(code, output)
}

func UpdateCategory(c *gin.Context){
	var UpdateCategoryRequestDTO requestsDTO.UpdateCategoryRequestDTO

	if err := c.ShouldBindJSON(&UpdateCategoryRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.UpdateCategory(UpdateCategoryRequestDTO)
	c.JSON(code, output)
}

func DeleteCategory(c *gin.Context){
	var DeleteCategoryRequestDTO requestsDTO.DeleteCategoryRequestDTO

	if err := c.ShouldBindJSON(&DeleteCategoryRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteCategory(DeleteCategoryRequestDTO)
	c.JSON(code, output)
}

func CategoryService(router *gin.RouterGroup) {
	router.GET("/categories", GetAllCategories)
	router.GET("/category/:id", GetCategory)
	router.POST("/category/create", CreateCategory)
	router.POST("/category/update", UpdateCategory)
	router.POST("/category/delete", DeleteCategory)
}