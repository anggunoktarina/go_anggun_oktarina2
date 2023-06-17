package payload

type CreateCategoryRequest struct {
	Name 	string 		`json:"name" gorm:"unique"`
	UserID uint `json:"user_id"`
}

type CreateCategoryResponse struct {
	CategoryID uint `json:"category_id"`
}

type UpdateCategoryRequest struct {
	Name 	string 		`json:"name" gorm:"unique"`
	CategoryID  uint      `json:"category_id"`
}

type UpdateCategoryResponse struct {
	CategoryID uint `json:"category"`
}

type GetCategoryResponse struct {
	CategoryID uint   `json:"category_id"`
	Name 	string 		`json:"name" gorm:"unique"`
}

type GetCategorysResponse struct {
	Categorys []GetCategoryResponse `json:"categorys"`
}
