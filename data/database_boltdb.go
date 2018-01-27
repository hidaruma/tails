package data

import (
	"database/sql"
)

type Sqlite struct {
	Path string
	db *sql.DB
}

func (s *Sqlite) DSN(path string) {
	s.Path =
}

func (s *Sqlite) Setup() error {

}