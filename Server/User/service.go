package user

import (
	"assessment-bootcamp/entity"
)

type service interface {
	UserGetAll() ([]entity.User, error)
	UserCreated(userCreate entity.User) (entity.User, error)
	UserbyID(UserbyID string) (entity.User, error)
	UserFindEmail(email string) (entity.User, error)
	UserDelete(UserDel entity.User) (entity.User, error)
}
