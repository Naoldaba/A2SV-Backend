package infrastructure

import (
	domain "task_manager_api/Domain"

	"golang.org/x/crypto/bcrypt"
)

type HashPassword func(user *domain.User) (string, error)
type ComparePassword func(user, existingUser *domain.User) bool

var PasswordHasher HashPassword = func(user *domain.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

var PasswordComparator ComparePassword = func(inputUser, storedUser *domain.User) bool {
	if inputUser == nil || storedUser == nil {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(inputUser.Password))
	return err == nil
}
