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
)

func GetAllUser(c *gin.Context){
	var GetAllUsersRequestDTO requestsDTO.GetAllUsersRequestDTO
	GetAllUsersRequestDTO.Page, GetAllUsersRequestDTO.Limit, GetAllUsersRequestDTO.OrderBy, GetAllUsersRequestDTO.OrderType = utils.PaginationHandler(GetAllUsersRequestDTO.Page, GetAllUsersRequestDTO.Limit, GetAllUsersRequestDTO.OrderBy, GetAllUsersRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllUsersRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllUser(GetAllUsersRequestDTO)
	c.JSON(code, output)
}

func GetUser(c *gin.Context){
	var GetUserRequestDTO requestsDTO.GetUserRequestDTO
	if err := c.ShouldBindWith(&GetUserRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetUser(GetUserRequestDTO)
	c.JSON(code, output)
}

func RegisterUser(c *gin.Context){
	var RegisterUserRequestDTO requestsDTO.RegisterUserRequestDTO

	if err := c.ShouldBindJSON(&RegisterUserRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.RegisterUser(RegisterUserRequestDTO)
	c.JSON(code, output)
}

func LoginUser(c *gin.Context){
	var LoginUserRequestDTO requestsDTO.LoginUserRequestDTO
	if err := c.ShouldBindJSON(&LoginUserRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code , output := helpers.LoginUser(LoginUserRequestDTO)
	c.JSON(code, output)
}

func UserService(router *gin.RouterGroup) {
	router.POST("/users/register", RegisterUser)
	router.POST("/users/login", LoginUser)
}
