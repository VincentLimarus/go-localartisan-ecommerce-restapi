package requestsDTO

type GetAllCategoriesRequestDTO struct {
	Page      int    `json:"page" form:"page" binding:"omitempty"`
	Limit     int    `json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetCategoryRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type CreateCategoryRequestDTO struct {
	Name      string `json:"name" form:"name" binding:"required"`
	Image     string `json:"image" form:"image" binding:"omitempty"`
	CreatedBy string `json:"created_by" form:"created_by" binding:"omitempty"`
	IsActive  bool   `json:"is_active" form:"is_active" binding:"omitempty"`
}

type UpdateCategoryRequestDTO struct {
	ID        string `json:"id" form:"id" binding:"required"`
	Name      string `json:"name" form:"name" binding:"omitempty"`
	Image     string `json:"image" form:"image" binding:"omitempty"`
	UpdatedBy string `json:"updated_by" form:"updated_by" binding:"omitempty"`
	IsActive  bool   `json:"is_active" form:"is_active" binding:"omitempty"`
}

type DeleteCategoryRequestDTO struct {
	// Only admin can delete.
	ID string `json:"id" form:"id" binding:"required"`
}