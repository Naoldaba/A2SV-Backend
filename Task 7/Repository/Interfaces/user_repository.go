package interfaces

import domain "task_manager_api/Domain"

type IUserRepository interface {
	GetUser(email string) (*domain.User, error)
	Register(user *domain.User) error
}