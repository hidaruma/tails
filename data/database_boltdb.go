package data

import (
	"github.com/boltdb/bolt"

	"os"
	"log"
	"path/filepath"
	"encoding/gob"
	"bytes"

	"go/ast"
	"fmt"
	"encoding/binary"
)

type BoltDB struct {
	Path string
	db *bolt.DB
}

func (bdb *BoltDB) DSN(path string) {
	bdb.Path = path
}

func (bdb *BoltDB) InitDB() error{
	_, err:= os.Stat(bdb.Path)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return err
	}
}

func (bdb *BoltDB) Open() error {
	db, err := bolt.Open(bdb.Path, 0600, nil)
	if err != nil {
		return err
	}

	bdb.db = db
	return nil
}

func (bdb *BoltDB) Close() error {
	return bdb.db.Close()
}

func (bdb *BoltDB) Cover{}

func (bdb *BoltDB) AllUsers() ([]*User, error) {
	users := make([]*User, 0)
	tmps, err := bdb.iterate([]byte("users"))
	if err != nil {
		return nil, err
	}
	for _, v := range tmps {
		encoded := v.([]byte)
		user, err := decodeUser(encoded)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (bdb *BoltDB) SaveUser(u *User) error {
	key := itob(u.ID)
	val, err := encode(u)
	if err != nil {
		return err
	}
	err = bdb.store([]byte("users"), key, val)
	if err != nil {
		return err
	}
	return nil
}

func (bdb *BoltDB) DeleteUser(u *User) error {
	key := itob(u.ID)
	return bdb.delete([]byte("users"), key)
}

func (bdb *BoltDB) LoadUser(u *User) error {
	if u.ID != 0 {
		key := itob(u.ID)
		if _, err := bdb.retrieve([]byte("users"), key) {

		}
	}
}

func (bdb *BoltDB) UpdateUser(u *User) error {
	return bdb.SaveUser(u)
}

func (bdb *BoltDB) userQuery(query string, args ...interface{}) ([]User, error) {
	switch query {
	case "all":
	case "delete":
	case "update":
	case "save":
	}
}


func (bdb *BoltDB) store(name []byte,key []byte, val []byte) error {
	err := bdb.db.Update(func(tx *bolt.Tx) error {
		bktCreated, err := tx.CreateBucketIfNotExists(name)
		if err != nil {
			return err
		}
		if bktCreated != nil {
			err = bktCreated.Put(key, val)
			if err != nil {
				return err
			}
			return nil
		} else {
			bktExists := tx.Bucket(name)
			err = bktExists.Put(key, val)
			if err != nil {
				return err
			}
			return nil
		}
	})
	return err
}

func (bdb *BoltDB) retrieve (name []byte, key []byte) ([]byte, error) {
	var result []byte
	err := bdb.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(name)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", name)
		}
		val := bucket.Get(key)
		result = val
		return nil
	})
	return result, err
}

func (bdb *BoltDB) retrieveAll(name []byte) () {}

func (bdb *BoltDB) iterate(name []byte) ([]interface{}, error) {
	var results []interface{}
	err := bdb.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(name)
		bucket.ForEach(func(key, val []byte) error {
			results = append(results, val)
			return nil
		})

	})
	return results, err
}

func (bdb *BoltDB) delete (name []byte, key []byte) error {
	err := bdb.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(name)
		err := bucket.Delete(key)
		return err
	})
	return err
}

func encode(arg interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	encoder := gob.NewEncoder(buf)
	err := encoder.Encode(arg)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decodeUser(arg []byte) (*User, error) {
	buf := bytes.NewBuffer(arg)
	var user *User
	decoder := gob.NewDecoder(buf)
	err := decoder.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func itob(v int) []byte {
	b := make([]byte, 0)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}