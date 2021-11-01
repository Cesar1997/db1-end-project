package main

import (
	"github.com/Cesar1997/db1-end-project/db"
	"github.com/Cesar1997/db1-end-project/structures"
)

func createUser(userRequest structures.User) error {
	err := db.RegisterUser(userRequest)
	if err != nil {
		return err
	}

	return nil
}

func loginUser(userRequest structures.User) (user structures.User, err error) {
	user, err = db.Login(userRequest)
	if err != nil {
		return user, err
	}
	return user, err
}
