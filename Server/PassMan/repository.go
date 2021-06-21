package passman

import (
	"assessment-bootcamp/entity"
)

type Repository interface {
	PassGetAll() ([]entity.PassMan, error)
	PassCreated(PassmanC entity.PassMan) (entity.PassMan, error)
	PassDelete()
	UserbyID(userID string) (entity.PassMan, error)
}
