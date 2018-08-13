package Proxy

import (
	"errors"
	"fmt"
)

type User struct {
	ID int32
}
type UserList []User

type UserListProxy struct {
	SomeDataBase        *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

func (t *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*t); i++ {
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %s could not be found\n", id)
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackSize {

	}

}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.LastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomeDataBase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	return User{}, errors.New("not implemented yet")
}
