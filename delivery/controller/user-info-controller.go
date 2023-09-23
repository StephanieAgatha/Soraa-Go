package controller

import (
	"github.com/StephanieAgatha/Soraa-Go/delivery/middleware"
	"github.com/StephanieAgatha/Soraa-Go/model"
	"github.com/StephanieAgatha/Soraa-Go/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUC usecase.UserUsecase
	gin    *gin.Engine
}

func (u UserController) CreateUserHandler(c *gin.Context) {
	//bind json
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	//usecase logic
	if err := u.UserUC.CreateUser(user); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Successfully Register"})
}

// create route method
func (u UserController) Route() {
	u.gin.POST("/new", middleware.AuthMiddleware(), u.CreateUserHandler)
}

func NewUserController(uc usecase.UserUsecase, g *gin.Engine) *UserController {
	return &UserController{
		UserUC: uc,
		gin:    g,
	}
}
