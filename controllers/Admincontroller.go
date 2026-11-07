package controllers

import (
	"net/http"

	"github.com/msdrepo/folcon/dto"
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

	profileData := adminData.EmailId + " " + adminData.Password

	response := res.SuccessResponse(http.StatusOK, "admin login successfully", profileData, nil)

	c.JSON(200, response)

}
