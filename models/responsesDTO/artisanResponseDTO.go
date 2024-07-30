package responsesDTO

type ArtisansResponseDTO struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	ShopName    string `json:"shop_name"`
	ShopAddress string `json:"shop_address"`
	Description string `json:"description"`
	ShopBanner  string `json:"shop_banner"`
	Rating      string `json:"rating"`
	IsActive    bool   `json:"is_active"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`

	User UserResponseDTO `json:"user"`
}