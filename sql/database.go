package sql

import (
	"database/sql"
)

type Database interface {
	Table(name string) Database

	ID(id any) Database
	In(string, ...any) Database
	Where(string, ...any) Database
	Columns(...string) Database
	AllCols() Database
	MustCols(...string) Database
	ShowSQL(showSQL bool) Database

	FindOne(document any, filter ...any) (bool, error)
	FindMany(documents any, filter ...any) error

	InsertOne(document any) (id any, err error)
	InsertMany(documents []any) ([]any, error)

	UpdateOne(document any) error

	DeleteOne(filter ...any) error

	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)

	Sync(...any) error

	Close() error
}
