package controllers

import (
	"net/http"

	res "github.com/msdrepo/folcon/responses"
	folcon "github.com/msdrepo/framework"
)

func AdminLogin(c *folcon.Context) {

	response := res.SuccessResponse(http.StatusOK, "login successfully", nil, nil)

	c.JSON(200, response)

}

func GetAdminProfile(c *folcon.Context) {

	userName := c.Query("user_name")
	role := c.Query("role")

	var profileData struct {
		UserName string `json:"user_name"`
		Role     string `json:"role"`
	}

	profileData.UserName = userName
	profileData.Role = role

	response := res.SuccessResponse(http.StatusOK, "profile fetched successfully", profileData, nil)

	c.JSON(200, response)

}
