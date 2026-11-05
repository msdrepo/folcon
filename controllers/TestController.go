package controllers

import (
	"net/http"

	res "github.com/msdrepo/folcon/responses"

	folcon "github.com/msdrepo/framework"
)

func Greet(c *folcon.Context) {

	response := res.SuccessResponse(http.StatusOK, "hello world", nil, nil)

	c.JSON(200, response)

}

func Login(c *folcon.Context) {

	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginRequest

	if err := c.BindJSON(&req); err != nil {
		res.ErrorResponse(http.StatusBadRequest, "login failed: invalid JSON or missing fields", err.Error())
		return
	}

	data := "test data"

	response := res.SuccessResponse(http.StatusOK, "login successfully", data, nil)

	c.JSON(200, response)

}

func TestCall(c *folcon.Context) {

	response := res.SuccessResponse(http.StatusOK, "hello", nil, nil)

	c.JSON(200, response)

}

func UserCall(c *folcon.Context) {

	name := c.Query("name")
	age := c.Query("age")

	response := res.SuccessResponse(http.StatusOK, "get data", "name "+name+" age "+age, nil)

	c.JSON(200, response)

}
