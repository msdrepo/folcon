package controllers

import (
	"net/http"

	"github.com/msdrepo/folcon/dto"
	"github.com/msdrepo/folcon/models"
	res "github.com/msdrepo/folcon/responses"
	"github.com/msdrepo/folcon/validations"
	folcon "github.com/msdrepo/framework"
)

func AdminLogin(c *folcon.Context) {

	var input validations.AdminLoginRequest

	if validationErrors := c.ValidateRequired(&input); len(validationErrors) > 0 {
		response := res.ErrorResponse(http.StatusUnprocessableEntity, "Validation failed", validationErrors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	adminData := dto.AdminLoginDTO{}

	adminData.EmailId = input.EmailId
	adminData.Password = input.Password

	token, err := models.AdminLoginCheck(adminData)

	if err != nil {

		response := res.ErrorResponse(http.StatusBadRequest, "email or password is incorrect", nil)

		c.JSON(200, response)
		return

	}

	response := res.SuccessResponse(http.StatusOK, "admin login successfully", nil, token)

	c.JSON(200, response)

}
