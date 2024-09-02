package requestsDTO

type GetAllUsersRequestDTO struct {
	Page      int    `json:"page" form:"page" binding:"omitempty"`
	Limit     int    `json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetUserRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type RegisterUserRequestDTO struct {
	Name            string `json:"name" form:"name" binding:"required"`
	Email           string `json:"email" form:"email" binding:"required"`
	Address         string `json:"address" form:"address" binding:"omitempty"`
	PhoneNumber     string `json:"phone_number" form:"phone_number" binding:"omitempty"`
	Password        string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
	IsActive        bool   `json:"is_active" form:"is_active" binding:"required"`
	CreatedBy       string `json:"created_by" form:"created_by" binding:"omitempty"`
}

type LoginUserRequestDTO struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UpdateUserRequestDTO struct {
	ID          string `json:"id" form:"id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"omitempty"`
	Email       string `json:"email" form:"email" binding:"omitempty"`
	PhoneNumber string `json:"phone_number" form:"phone_number" binding:"omitempty"`
	Address     string `json:"address" form:"address" binding:"omitempty"`
	IsActive    bool   `json:"is_active" form:"is_active" binding:"omitempty"`
	UpdatedBy   string `json:"updated_by" form:"updated_by" binding:"omitempty"`
}

type DeleteUserRequestDTO struct {
	ID              string `json:"id" form:"id" binding:"required"`
	Password        string `json:"password" form:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
}

type ChangePasswordRequestDTO struct {
	ID              string `json:"id" form:"id" binding:"required"`
	OldPassword     string `json:"old_password" form:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" form:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password" binding:"required"`
	UpdatedBy       string `json:"updated_by" form:"updated_by" binding:"omitempty"`
}
