package interfaces

import "go/tutorial/entity"

type UserInterface interface {
	FindAll() []entity.User
	FindById(id int) entity.User
	Save(u entity.User) entity.User
	DeleteById(id int) bool
	UpdateById(id int, u entity.User) entity.User
}
