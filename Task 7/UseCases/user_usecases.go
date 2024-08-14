package usecases

import (
	"errors"
	"log"
	"task_manager_api/Domain"
	infrastructure "task_manager_api/Infrastructure"
	"task_manager_api/Repository/Interfaces"
)

type UserUseCase struct {
	userRepo interfaces.IUserRepository
	hasher infrastructure.HashPassword
}

func NewUserUseCase(userRepo interfaces.IUserRepository,hasher infrastructure.HashPassword ) *UserUseCase{
	return &UserUseCase{
		userRepo:userRepo,
		hasher: hasher,
	}
}

func (uc *UserUseCase) Register(user *domain.User) error {
	existingUser, _ := uc.userRepo.GetUser(user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}
	hashedPassword, err := uc.hasher(user)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = hashedPassword
	return uc.userRepo.Register(user)
}

func (uc *UserUseCase) GetUserByEmail(email string) (*domain.User, error) {
	user, err := uc.userRepo.GetUser(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}