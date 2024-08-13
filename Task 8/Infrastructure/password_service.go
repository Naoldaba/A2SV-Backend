package infrastructure

import (
	domain "task_manager_api/Domain"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(user *domain.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}