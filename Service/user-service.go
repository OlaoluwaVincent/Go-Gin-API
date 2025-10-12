package service

import (
	"go/tutorial/entity"
)

type UserService struct {
	users []entity.User
}

func (u *UserService) FindAll() []entity.User {
	return u.users
}

func (u *UserService) FindById(id int) entity.User {
	for i, user := range u.users {
		if user.Id == id {
			return u.users[i]
		}
	}
	return entity.User{}
}

func (u *UserService) Save(user entity.User) entity.User {
	u.users = append(u.users, user)
	return user
}

func (u *UserService) DeleteById(id int) bool {
	for i, user := range u.users {
		if user.Id == id {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return true
		}
	}
	return false
}

func (u *UserService) UpdateById(id int, user entity.User) entity.User {
	for i, v := range u.users {
		if v.Id == id {
			u.users[i] = user
			return u.users[i]
		}
	}
	return entity.User{}
}

// Call this function to Instantiate the VideoService
// After instantiate, you can call the methods on the returned object
func NewUserService() *UserService {
	return &UserService{}
}
