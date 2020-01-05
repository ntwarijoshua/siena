package Handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ntwarijoshua/siena/internal/models"
	"github.com/ntwarijoshua/siena/internal/services"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CreateUserRequest struct {
	Email    string `json:"email" validate:"email,is_unique"`
	Password string `json:"password" validate:"min=6"`
	Names    string `json:"names" validate:"required"`
	DOB      string `json:"date_of_birth" validate:"required"`
}

type AuthenticateUserRequest struct {
	Email string 	`json:"email" validate:"email"`
	Password string `json:"password" validate:"required"`
}



func (app *App)CreateUser(c *gin.Context) {
	var (
		payload CreateUserRequest
		validationService = app.ServiceContainer.GetService("validationService").(*services.ValidationService)
		dateLayout = "2006-01-02"
		userService = app.ServiceContainer.GetService("userService").(*services.UserService)
		)


	if err := c.BindJSON(&payload); err != nil {
		app.Logger.Errorf("Error occurred while decoding user payload:", err)
		c.JSON(http.StatusBadRequest, map[string]string{"error": "failed parsing payload"})
		return
	}
	validate := validationService.GetValidator()
	if err := validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, validationService.GenerateValidationResponse(err))
		return
	}
	dob, err := time.Parse(dateLayout, payload.DOB)
	if err != nil {
		logger.Errorf("Error occurred while parsing user date of birth", err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error occurred": err})
		return
	}
	user := models.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	profile := models.Profile{
		Names:   payload.Names,
		TagLine: "",
		DOB:     dob,
	}
	newUser, err := userService.CreateUser(user, profile)

	if err != nil {
		logger.Errorf("Error occurred trying to create user", err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"errors occurred": err})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user created successfully",
		"data":    newUser,
	})
}

func (app *App) AuthenticateUser (c *gin.Context)  {
	var (
		payload AuthenticateUserRequest
		usersService = app.ServiceContainer.GetService("userService").(*services.UserService)
		validationService = app.ServiceContainer.GetService("validationService").(*services.ValidationService)
	)

	if err := c.BindJSON(&payload); err != nil {
		app.Logger.Errorf("Error occurred while decoding user payload:", err)
		c.JSON(400, map[string]string{"error": "failed parsing payload"})
		return
	}
	validate := validationService.GetValidator()
	if err := validate.Struct(payload); err != nil {
		c.JSON(http.StatusBadRequest, validationService.GenerateValidationResponse(err))
		return
	}
	user, err := usersService.GetUserByMail(payload.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]string{
			"error": fmt.Sprintf("User with this email not found %s", err),
		})
		return
	}
	token, err := usersService.GetJWTToken(user, payload.Password)
	if err != nil {
		app.Logger.Errorf("Error occurred while generating user token:", err)
		c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed generating token"})
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"message" : "Authenticate Successfully",
		"token":token,
	})
}




