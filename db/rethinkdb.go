package db

import (
	r "github.com/dancannon/gorethink"
	config "github.com/namKolo/shorturl/config"
)

// DB contains rethinkdb session
type DB struct {
	session *r.Session
}

// New return a rethinkdb session
func New(dbConfig *config.RethinkDB) (*DB, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  dbConfig.Host,
		Database: dbConfig.DB,
		Password: dbConfig.Password,
	})
	if err != nil {
		return nil, err
	}

	return &DB{session}, nil
}
