package validations

// request json field validation struct
type LoginRequest struct {
	EmailId  string `json:"email_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetProductsInput struct {
	SearchText string `json:"search_text"`
	PageNumber uint   `json:"page_number"`
	NumOfItems uint   `json:"num_of_items"`
}
