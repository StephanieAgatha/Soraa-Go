package controller

import (
	"github.com/StephanieAgatha/Soraa-Go/model"
	"github.com/StephanieAgatha/Soraa-Go/usecase"
	"github.com/gin-gonic/gin"
)

type UserCredentialController struct {
	userCredUc usecase.UserCredentialUsecase
	gin        *gin.Engine
}

func (u UserCredentialController) Register(c *gin.Context) {
	var userCred model.UserCredentials

	if err := c.ShouldBindJSON(&userCred); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad JSON Format"})
		return
	}

	if err := u.userCredUc.Register(userCred); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Message": "Successfully Register"})
}

func (u UserCredentialController) Login(c *gin.Context) {
	var userCred model.UserCredentials

	if err := c.ShouldBindJSON(&userCred); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "Bad JSON Format"})
		return
	}

	_, err := u.userCredUc.FindUserEMail(userCred.Email)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	userToken, err := u.userCredUc.Login(userCred)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"Data": userToken})
}

func (u UserCredentialController) Route() {
	authGroup := u.gin.Group("/auth")
	{
		authGroup.POST("/register", u.Register)
		authGroup.POST("/login", u.Login)
	}
}
func NewUserCredentialController(uc usecase.UserCredentialUsecase, g *gin.Engine) *UserCredentialController {
	return &UserCredentialController{
		userCredUc: uc,
		gin:        g,
	}
}
