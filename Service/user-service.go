package service

import (
	"go/tutorial/entity"
	"reflect"
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
	user.Id = len(u.users) + 1
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
			vValue := reflect.ValueOf(&v).Elem()
			uValue := reflect.ValueOf(user)

			for j := 0; j < vValue.NumField(); j++ {
				field := vValue.Field(j)
				newValue := uValue.Field(j)

				if !reflect.DeepEqual(newValue.Interface(), reflect.Zero(newValue.Type()).Interface()) {
					field.Set(newValue)
				}
			}

			u.users[i] = v
			return v
		}
	}
	return entity.User{}
}

// Call this function to Instantiate the VideoService
// After instantiate, you can call the methods on the returned object
func NewUserService() *UserService {
	return &UserService{}
}
