package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func StringToUUID(s string) uuid.UUID {
	id, _ := uuid.Parse(s)
	return id
}

func UUIDToString(id uuid.UUID) string {
	return id.String()
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func PageHandler(page int) int {
	if page == 0 {
		return 1
	}
	return page
}

func LimitHandler(limit int) int {
	if limit == 0 {
		return 10
	}
	return limit
}

func OrderByHandler(order_by string) string {
	if order_by == "" {
		return "ID"
	}
	return order_by
}

func OrderTypeHandler(order_type string) string {
	if order_type == "" {
		return "asc"
	}
	return order_type
}

func PaginationHandler(page int, limit int, order_by string, order_type string) (int, int, string, string) {
	page = PageHandler(page)
	limit = LimitHandler(limit)
	order_by = OrderByHandler(order_by)
	order_type = OrderTypeHandler(order_type)

	return page, limit, order_by, order_type
}