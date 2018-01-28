/*
Copyright (C) 2014 Matt Layher
Released under the MIT license
http://opensource.org/licenses/mit-license.php
 */
package data

import (
	"time"
	"golang.org/x/crypto/pbkdf2"
	"fmt"
	"crypto/sha1"
	"crypto/rand"
)

type Session struct {
	ID int
	UserID int
	Client string
	Expire int64
	Key string
}

func NewSession(userID int, password string, client string) (*Session, error) {

	session := &Session{
		UserID: userID,
		Client: client,
	}

	session.Expire = time.Now().Add(time.Duration(7 * 24 *time.Hour)).Unix()

	saltBuf := make([]byte, 16)
	if err := rand.Read(saltBuf); err != nil {
		return nil, err
	}
	salt1 := saltBuf
	session.Key = fmt.Sprintf("%x", pbkdf2.Key([]byte(password), salt1, 4096, 16,sha1.New))

	if err := session.Save(); err != nil {
		return nil, err
	}
	return session, nil
}

func (s *Session) Save() error {
	return DataBase.SaveSession(s)
}

func (s *Session) Delete() error {
	return DataBase.DeleteSession(s)
}

func (s *Session) Update() error {
	return DataBase.UpdateSession(s)
}

func (s *Session) Load() error {
	return DataBase.LoadSession(s)
}