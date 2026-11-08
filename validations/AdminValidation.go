package validations

type AdminLoginRequest struct {
	EmailId  string `json:"email_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
