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

func GetAllArtisans(c *gin.Context){
	var GetAllArtisansRequestDTO requestsDTO.GetAllArtisansRequestDTO
	GetAllArtisansRequestDTO.Page, GetAllArtisansRequestDTO.Limit, GetAllArtisansRequestDTO.OrderBy, GetAllArtisansRequestDTO.OrderType = utils.PaginationHandler(GetAllArtisansRequestDTO.Page, GetAllArtisansRequestDTO.Limit, GetAllArtisansRequestDTO.OrderBy, GetAllArtisansRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllArtisansRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllArtisans(GetAllArtisansRequestDTO)
	c.JSON(code, output)
}

func GetArtisan(c *gin.Context) {
	artisanID := c.Param("id")

	if _, err := uuid.Parse(artisanID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	
	code, output := helpers.GetArtisan(artisanID)
	c.JSON(code, output)
}

func RegisterArtisan(c *gin.Context){
	var RegisterArtisanRequestDTO requestsDTO.RegisterArtisanRequestDTO
	var UserInformation requestsDTO.UserInformation

	if err := c.ShouldBindJSON(&RegisterArtisanRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.RegisterArtisan(RegisterArtisanRequestDTO, UserInformation)
	c.JSON(code, output)
}

func UpdateArtisan(c *gin.Context){
	var UpdateArtisanRequestDTO requestsDTO.UpdateArtisanRequestDTO

	if err := c.ShouldBindJSON(&UpdateArtisanRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.UpdateArtisan(UpdateArtisanRequestDTO)
	c.JSON(code, output)
}

func DeleteArtisan(c *gin.Context){
	var DeleteArtisanRequestDTO requestsDTO.DeleteArtisanRequestDTO

	if err := c.ShouldBindJSON(&DeleteArtisanRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteArtisan(DeleteArtisanRequestDTO)
	c.JSON(code, output)
}

func BaseArtisanService(router *gin.RouterGroup) {
	router.GET("/artisans", GetAllArtisans)
	router.GET("/artisan/:id", GetArtisan)
}

func AuthArtisanService(router *gin.RouterGroup) {
	router.POST("/artisan/register", RegisterArtisan)
	router.POST("/artisan/update", UpdateArtisan)
	router.POST("/artisan/delete", DeleteArtisan)
}

