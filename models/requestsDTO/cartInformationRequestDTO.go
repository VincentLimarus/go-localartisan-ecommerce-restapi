package requestsDTO

import "github.com/google/uuid"

type GetAllCartInformationsRequestDTO struct {
	Page      int    `json:"page" form:"page" query:"page"`
	Limit     int    `json:"limit" form:"limit" query:"limit"`
	OrderBy   string `json:"orderBy" form:"orderBy" query:"orderBy"`
	OrderType string `json:"orderType" form:"orderType" query:"orderType"`
}

type GetAllCartInformationByCartIDRequestDTO struct {
	CartID uuid.UUID `json:"cart_id" form:"cart_id" binding:"required"`
}

type AddItemToCartRequestDTO struct {
	CartID       uuid.UUID `json:"cart_id" form:"cart_id" binding:"required"`
	ProductID    uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	Quantity     int       `json:"quantity" form:"quantity" binding:"required"`
	IsActive     bool      `json:"is_active" form:"is_active" binding:"required"`
	CreatedBy    string    `json:"created_by" form:"created_by" binding:"omitempty"`
}

type UpdateItemInCartRequestDTO struct {
	CartID       uuid.UUID `json:"cart_id" form:"cart_id" binding:"required"`
	ProductID    uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	Quantity     int       `json:"quantity" form:"quantity" binding:"omitempty"`
	IsActive     bool      `json:"is_active" form:"is_active" binding:"omitempty"`
	UpdatedBy    string    `json:"updated_by" form:"updated_by" binding:"omitempty"`
}

type DeleteItemInCartRequestDTO struct {
	CartID    uuid.UUID `json:"cart_id" form:"cart_id" binding:"required"`
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
}