package requestsDTO

import "github.com/google/uuid"

type GetAllCartInformationsRequestDTO struct {
	Page      int    `json:"page" form:"page" query:"page"`
	Limit     int    `json:"limit" form:"limit" query:"limit"`
	OrderBy   string `json:"orderBy" form:"orderBy" query:"orderBy"`
	OrderType string `json:"orderType" form:"orderType" query:"orderType"`
}

type GetAllCartInformationByCartIDRequestDTO struct {
	CartID uuid.UUID `json:"cartID" form:"cartID" binding:"required"`
}

type AddItemToCartRequestDTO struct {
	CartID       uuid.UUID `json:"cartID" form:"cartID" binding:"required"`
	ProductID    uuid.UUID `json:"productID" form:"productID" binding:"required"`
	Quantity     int       `json:"quantity" form:"quantity" binding:"required"`
	PriceAtOrder float64   `json:"priceAtOrder" form:"priceAtOrder" binding:"required"`
	IsActive     bool      `json:"isActive" form:"isActive" binding:"required"`
	CreatedBy    string    `json:"createdBy" form:"createdBy" binding:"omitempty"`
}

type UpdateItemInCartRequestDTO struct {
	CartID       uuid.UUID `json:"cartID" form:"cartID" binding:"required"`
	ProductID    uuid.UUID `json:"productID" form:"productID" binding:"required"`
	Quantity     int       `json:"quantity" form:"quantity" binding:"omitempty"`
	PriceAtOrder float64   `json:"priceAtOrder" form:"priceAtOrder" binding:"omitempty"`
	IsActive     bool      `json:"isActive" form:"isActive" binding:"omitempty"`
	UpdatedBy    string    `json:"updatedBy" form:"updatedBy" binding:"omitempty"`
}

type DeleteItemInCartRequestDTO struct {
	CartID    uuid.UUID `json:"cartID" form:"cartID" binding:"required"`
	ProductID uuid.UUID `json:"productID" form:"productID" binding:"required"`
}