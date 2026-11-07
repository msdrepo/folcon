package dto

type AdminLoginDTO struct {
	EmailId  string `json:"email_id"`
	Password string `json:"password"`
}

type AdminDataDTO struct {
	AdminId  uint   `json:"admin_id"`
	FullName string `json:"full_name"`
	EmailId  string `json:"email_id"`
	Password string `json:"password"`
}
