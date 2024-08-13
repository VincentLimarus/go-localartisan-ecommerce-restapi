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

func GetUserByID(c *gin.Context){
	userID := c.Param("id")

	if _, err := uuid.Parse(userID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetUser(userID)
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
	code , output, token := helpers.LoginUser(LoginUserRequestDTO)
	c.SetCookie("Authorization", token, 3600*12, "/", "localhost", true, true)
	c.SetSameSite(http.StatusOK)
	c.JSON(code, output)
}

func UpdateUser(c *gin.Context){
	var UpdateUserRequestDTO requestsDTO.UpdateUserRequestDTO
	if err := c.ShouldBindJSON(&UpdateUserRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.UpdateUser(UpdateUserRequestDTO)
	c.JSON(code, output)
}

func DeleteUser(c *gin.Context){
	var DeleteUserRequestDTO requestsDTO.DeleteUserRequestDTO
	if err := c.ShouldBindJSON(&DeleteUserRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteUser(DeleteUserRequestDTO)
	c.JSON(code, output)
}

func ChangePasswordUser(c *gin.Context){
	var ChangePasswordRequestDTO requestsDTO.ChangePasswordRequestDTO
	if err := c.ShouldBindJSON(&ChangePasswordRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.ChangePasswordUser(ChangePasswordRequestDTO)
	c.JSON(code, output)
}

func BaseUserService(router *gin.RouterGroup) {
	router.POST("/user/register", RegisterUser)
	router.POST("/user/login", LoginUser)
}

func AuthUserService(router *gin.RouterGroup) {
	router.GET("/users", GetAllUser)
	router.GET("/user/:id", GetUserByID)
	router.POST("/user/update", UpdateUser)
	router.POST("/user/delete", DeleteUser)
	router.POST("/user/change-password", ChangePasswordUser)
}

