package repository

import (
	"github.com/jemmycalak/calak_chatdate_postgre/src/moduls/user/model"
)

type UserRepository interface {
	Save(*model.User) error
	Upadate(string, *model.User) error
	Delete(string) error
	FindById(string) (*model.User, error)
	FindAll() (model.Users, error)
}
