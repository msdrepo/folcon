package controllers

import (
	"net/http"

	res "github.com/msdrepo/folcon/responses"
	folcon "github.com/msdrepo/framework"
)

func GetProfile(c *folcon.Context) {

	profileData := map[string]string{"name": "Augustin", "role": "Backend Developer"}

	response := res.SuccessResponse(http.StatusOK, "Profile fetched successfully", profileData, nil)

	c.JSON(200, response)

}
