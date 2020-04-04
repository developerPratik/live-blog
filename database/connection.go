package database

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"os"
)

type Db struct {
	Connection r.Connection
	Session    *r.Session
}

func (db *Db) CreateConnection() error {

	session, err := r.Connect(r.ConnectOpts{
		Address:  os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		return err
	}

	db.Session = session
	return nil

}
func Get() *Db {
	return &Db{}
}
