package models

import (
	"database/sql"
	"errors"

	"github.com/msdrepo/folcon/dto"
	folcon "github.com/msdrepo/framework"
)

func AdminLoginCheck(data dto.AdminLoginDTO) (string, error) {

	var token string
	var err error

	var adminData dto.AdminDataDTO

	query := `select id AS admin_id, full_name, email_id, password AS hashed_password 
	from admins where email_id = ?`

	err = folcon.DB.QueryRow(query, data.EmailId).Scan(
		&adminData.AdminId,
		&adminData.FullName,
		&adminData.EmailId,
		&adminData.HashedPassword,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("admin data not found")
		}
		return "", err
	}

	if err = folcon.VerifyPassword(data.Password, adminData.HashedPassword); err != nil {

		return "", err
	}

	token = "agaagaagaagaagaagaagaagaagaaga"

	return token, nil

}
