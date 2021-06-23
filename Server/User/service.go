package user

import (
	"assessment-bootcamp/Server/entity"
	"assessment-bootcamp/Server/helper"
	"errors"
	"fmt"

	"time"
)

type UserService interface {
	SaveNewUser(user entity.UserInput) (entity.User, error)
	LoginUser(input entity.UserInput) (entity.User, error)
	UpdateUserByID(id string, dataInput entity.UpdateUserInput) (entity.User, error)
	GetUserByID(id string) (entity.User, error)
}

type userService struct {
	repository UserRepository
}

func UserNewService(repository UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) SaveNewUser(user entity.UserInput) (entity.User, error) {
	var newUser = entity.User{
		FullName: user.FullName,
		Address:  user.Address,
		Email:    user.Email,
		Password: user.Password,
	}

	createUser, err := s.repository.CreateUser(newUser)

	if err != nil {
		return createUser, err
	}

	return createUser, nil
}

func (s *userService) LoginUser(input entity.UserInput) (entity.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %d not found", user.ID)
		return user, errors.New(newError)
	}

	if input.Password != user.Password {
		return user, errors.New("invalid password")
	}

	return user, nil
}

func (s *userService) UpdateUserByID(id string, dataInput entity.UpdateUserInput) (entity.User, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.User{}, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return entity.User{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return entity.User{}, errors.New(newError)
	}

	if dataInput.FullName != "" || len(dataInput.FullName) != 0 {
		dataUpdate["full_name"] = dataInput.FullName
	}
	if dataInput.Address != "" || len(dataInput.Address) != 0 {
		dataUpdate["address"] = dataInput.Address
	}
	dataUpdate["updated_at"] = time.Now()

	userUpdated, err := s.repository.UpdateUser(id, dataUpdate)

	if err != nil {
		return entity.User{}, err
	}

	return userUpdated, nil
}

func (s *userService) GetUserByID(id string) (entity.User, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.User{}, err
	}

	user, err := s.repository.GetOneUser(id)

	if err != nil {
		return entity.User{}, err
	}

	if user.ID == 0 {
		newError := fmt.Sprintf("user id %s is not found", id)
		return entity.User{}, errors.New(newError)
	}

	return user, nil
}
