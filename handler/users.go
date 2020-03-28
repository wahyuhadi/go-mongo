package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"framework/utils"
	"framework/models"
	"framework/forms"
	"framework/services"
	"framework/helpers"
)


type UserHandler struct {
	Log     *logrus.Logger
	Environ utils.Environ
}



func NewUsersHandler(logger *logrus.Logger, environ utils.Environ) *UserHandler {
	return &UserHandler{Log: logger, Environ: environ}
}



var (
	userModel = new(models.UserModel)
)

// Controller user signup
func (userHandler *UserHandler) SignUp (c *gin.Context) {
	var data forms.SignupUserCommand
	if c.BindJSON(&data) != nil {
		// specified response
		c.JSON(406, gin.H{"message": "Provide relevant fields"})
		// abort the request
		c.Abort()
		// return nothing
		return
	}

	checkEmail , _:= userModel.GetUserByEmail(data.Email)
	if checkEmail.Email != "" {
		c.JSON(403, gin.H{"message": "Email is already in use"})
		c.Abort()
		return
	}

	err := userModel.Signup(data)
	// Check if there was an error when saving user
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an account"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "New user account registered"})
}

func (userHandler *UserHandler) Login(c *gin.Context) {
	var data forms.LoginUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	result, err := userModel.GetUserByEmail(data.Email)

	if result.Email == "" {
		c.JSON(404, gin.H{"message": "User account was not found"})
		c.Abort()
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem logging into your account"})
		c.Abort()
		return
	}

	hashedPassword := []byte(result.Password)
	// Get the password provided in the request.body
	password := []byte(data.Password)

	err = helpers.PasswordCompare(password, hashedPassword)

	if err != nil {
		c.JSON(403, gin.H{"message": "Invalid user credentials"})
		c.Abort()
		return
	}

	jwtToken, err2 := services.GenerateToken(data.Email)

	// If we fail to generate token for access
	if err2 != nil {
		c.JSON(403, gin.H{"message": "There was a problem logging you in, try again later"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "Log in success", "token": jwtToken})

}



func (userHandler *UserHandler) GetDetails(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(403, gin.H{"message": "token not found"})
		c.Abort()
		return
	}

	userData, err := services.DecodeToken(token)
	if err != nil {
		c.JSON(403, gin.H{"message": "Invalid token"})
		c.Abort()
		return
	}

	isData, _ := userModel.GetUserDetails(userData)
	c.JSON(200, gin.H{"message": "User Details", "data": isData})

}


func (userHandler *UserHandler) Setup(r *gin.Engine) {
	r.POST("/users", userHandler.SignUp)
	r.POST("/users/auth", userHandler.Login)
	r.GET("/users", userHandler.GetDetails)
}