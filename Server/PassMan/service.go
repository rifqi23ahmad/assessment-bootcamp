package password

import (
	"assessment-bootcamp/Server/entity"
	"assessment-bootcamp/Server/helper"
	"errors"
	"fmt"
	"time"
)

type PassManService interface {
	SCreateNewPass(userID int, password entity.InputPassword) (entity.PassMan, error)
	SGetPassID(id string) (entity.PassMan, error)
	SUpdatePass(id string, dataInput entity.InputPassword) (entity.PassMan, error)
	SDelPassID(id string) (interface{}, error)
}

type passwordService struct {
	repository PassManRepository
}

func NewService(repository PassManRepository) *passwordService {
	return &passwordService{repository}
}

func (s *passwordService) SCreateNewPass(userID int, password entity.InputPassword) (entity.PassMan, error) {
	var newPass = entity.PassMan{
		UserID:   userID,
		Website:  password.Website,
		Password: password.Password,
	}

	createPassword, err := s.repository.RepCreatePass(newPass)

	if err != nil {
		return createPassword, err
	}

	return createPassword, nil
}

func (s *passwordService) SGetPassID(id string) (entity.PassMan, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.PassMan{}, err
	}

	password, err := s.repository.RepGetPass(id)

	if err != nil {
		return entity.PassMan{}, err
	}

	if password.ID == 0 {
		newError := fmt.Sprintf("password id %s is not found", id)
		return entity.PassMan{}, errors.New(newError)
	}

	return password, nil
}

func (s *passwordService) SUpdatePass(id string, dataInput entity.InputPassword) (entity.PassMan, error) {
	var dataUpdate = map[string]interface{}{}

	if err := helper.ValidateIDNumber(id); err != nil {
		return entity.PassMan{}, err
	}

	password, err := s.repository.RepGetPass(id)

	if err != nil {
		return entity.PassMan{}, err
	}

	if password.ID == 0 {
		newError := fmt.Sprintf("password id %s is not found", id)
		return entity.PassMan{}, errors.New(newError)
	}

	if dataInput.Website != "" || len(dataInput.Website) != 0 {
		dataUpdate["website"] = dataInput.Website
	}
	if dataInput.Password != "" || len(dataInput.Password) != 0 {
		dataUpdate["password"] = dataInput.Password
	}

	dataUpdate["updated_at"] = time.Now()

	passwordUpdated, err := s.repository.RepUpdatePass(id, dataUpdate)

	if err != nil {
		return entity.PassMan{}, err
	}

	return passwordUpdated, nil
}

func (s *passwordService) SDelPassID(id string) (interface{}, error) {
	if err := helper.ValidateIDNumber(id); err != nil {
		return nil, err
	}

	password, err := s.repository.RepGetPass(id)

	if err != nil {
		return nil, err
	}

	if password.ID == 0 {
		newError := fmt.Sprintf("password id %s is not found", id)
		return nil, errors.New(newError)
	}

	status, err := s.repository.RepDelPass(id)

	if err != nil {
		panic(err)
	}

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("password id %s success deleted", id)

	formatDelete := FormatDelete(msg)

	return formatDelete, nil
}
