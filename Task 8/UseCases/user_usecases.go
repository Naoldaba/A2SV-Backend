package usecases

import (
	"errors"
	"log"
	domain "task_manager_api/Domain"
	infrastructure "task_manager_api/Infrastructure"
	interfaces "task_manager_api/Repository/Interfaces"
)

type IUserUseCase interface {
	Register(user *domain.User) error
	GetUser(email string) (*domain.User, error)
	PromoteUserByID(id string, promoter *domain.User) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
}

type UserUseCase struct {
	userRepo interfaces.IUserRepository
	hasher   infrastructure.HashPassword
}

func NewUserUseCase(userRepo interfaces.IUserRepository, hasher infrastructure.HashPassword) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		hasher:   hasher,
	}
}

func (uc *UserUseCase) Register(user *domain.User) error {
	users, err := uc.userRepo.GetAllUsers()
	if err != nil {
		return errors.New("error while fetching")
	}
	if len(users) == 0{
		user.Role = "ADMIN"
	} else {
		user.Role = "USER"
		existingUser, _ := uc.userRepo.GetUser(user.Email)
		if existingUser != nil {
			return errors.New("user already exists")
		}
	}

	if user.Password == "" {
		return errors.New("password cannot be empty")
	}
	hashedPassword, err := uc.hasher(user)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = hashedPassword
	return uc.userRepo.Register(user)
}

func (uc *UserUseCase) GetUser(email string) (*domain.User, error) {
	user, err := uc.userRepo.GetUser(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) PromoteUserByID(id string, promoter *domain.User) (*domain.User, error) {
    if promoter.Role != "ADMIN" {
        return nil, errors.New("only admins can promote users")
    }
    updatedUser, err := uc.userRepo.PromoteUser(id)
    if err != nil {
        return nil, err
    }
    return updatedUser, nil
}

func (uc *UserUseCase) GetAllUsers() ([]*domain.User, error) {
    return uc.userRepo.GetAllUsers()
}