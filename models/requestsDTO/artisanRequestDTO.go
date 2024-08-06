package requestsDTO

import "github.com/google/uuid"

type GetAllArtisansRequestDTO struct {
	Page      int    `json:"page" form:"page" binding:"omitempty"`
	Limit     int    `json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetArtisansRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type RegisterArtisanRequestDTO struct {
	UserID      uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	ShopName    string    `json:"shop_name" form:"shop_name" binding:"required"`
	ShopAddress string    `json:"shop_address" form:"shop_address" binding:"required"`
	Description string    `json:"description" form:"description" binding:"omitempty"`
	ShopBanner  string    `json:"shop_banner" form:"shop_banner" binding:"omitempty"`
	IsActive    bool      `json:"is_active" form:"is_active" binding:"omitempty"`
	CreatedBy   string    `json:"created_by" form:"created_by" binding:"omitempty"`
}

type UpdateArtisanRequestDTO struct {
	ID          string `json:"id" form:"id" binding:"required"`
	UserID      uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	ShopName    string `json:"shop_name" form:"shop_name" binding:"omitempty"`
	ShopAddress string `json:"shop_address" form:"shop_address" binding:"omitempty"`
	Description string `json:"description" form:"description" binding:"omitempty"`
	ShopBanner  string `json:"shop_banner" form:"shop_banner" binding:"omitempty"`
	IsActive    bool   `json:"is_active" form:"is_active" binding:"omitempty"`
	UpdatedBy   string `json:"updated_by" form:"updated_by" binding:"omitempty"`
}

type DeleteArtisanRequestDTO struct {
	ID              string `json:"id" form:"id" binding:"required"`
	UserID 			uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	Password 		string `json:"password" form:"password" binding:"required"` // Password User
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}
