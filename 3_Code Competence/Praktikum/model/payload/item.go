package payload

type CreateItemRequest struct {
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`
	Stock       uint      `json:"stock"`
	Price       uint      `json:"price"`
	CategoryID  uint      `json:"category_id"`
}

type CreateItemResponse struct {
	ItemID uint `json:"item_id"`
}

type UpdateItemRequest struct {
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`
	Stock       uint      `json:"stock"`
	Price       uint      `json:"price"`
	CategoryID  uint      `json:"category_id"`
}

type UpdateItemResponse struct {
	ItemID uint `json:"item"`
}

type GetItemResponse struct {
	ItemID uint   `json:"item_id"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description"`
	Stock       uint      `json:"stock"`
	Price       uint      `json:"price"`
	CategoryID  uint      `json:"category_id"`
}

type GetItemsResponse struct {
	Items []GetItemResponse `json:"items"`
}
