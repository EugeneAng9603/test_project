package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLSaver struct {
	db *sql.DB
}

func (s *MySQLSaver) Save(data string) error {
	_, err := s.db.Exec("INSERT INTO table_name (data) VALUES (?)", data)
	return err
}

func NewMySQLSaver(dsn string) (*MySQLSaver, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &MySQLSaver{db: db}, nil
}
