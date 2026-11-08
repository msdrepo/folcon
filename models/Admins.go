package models

import (
	"github.com/msdrepo/folcon/dto"
	folcon "github.com/msdrepo/framework"
)

func AdminLoginCheck(data dto.AdminLoginDTO) (string, error) {

	var token string
	var err error

	token = "agaagaagaagaagaagaagaagaagaaga"

	HashedPassword := "$2y$10$3NA2pPYllJC4RF4VLRIyBuQHoq7i1tsEakc8900XTPXhSzVtrTcHW"

	if err := folcon.VerifyPassword(data.Password, HashedPassword); err != nil {

		return "", err
	}

	return token, nil

}
