/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */

package data

import (
	"golang.org/x/crypto/bcrypt"
	"encoding/hex"

)

const (
	GuestAccount = iota
	UserAccount
	AdminAccount
)

type User struct {
	ID int
	Name string
	Password string
	Icon string
	Email string
	AccountPermission int
}

func NewUser(name string, password string, icon string, email string, accountpermission int) (*User, error) {
	user := &User{
		Name: name,
		Email: email,
		AccountPermission: accountpermission,
	}
	if err := user.SetPassword(password); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) SetPassword(password string) error {
	converted, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return err
	}
	u.Password = hex.EncodeToString(converted)
	return nil
}

func (u *User) Save() error {
	return DataBase.SaveUser(u)
}

func (u *User) Update() error {
	return DataBase.UpdateUser(u)
}

func (u *User) Delete() error {
	return DataBase.DeleteUser(u)
}

func (u *User) Load() error {
	return DataBase.LoadUser(u)
}

func(u *User) setDefaultIcon() error {

}