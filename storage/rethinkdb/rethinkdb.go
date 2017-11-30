package rethinkdb

import (
	r "github.com/dancannon/gorethink"
	config "github.com/namKolo/shorturl/config"
	storage "github.com/namKolo/shorturl/storage"
)

// DB contains rethinkdb session
type rethinkdb struct {
	session *r.Session
}

// New return a rethinkdb session
func New(dbConfig *config.RethinkDB) (storage.Storage, error) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  dbConfig.Host,
		Database: dbConfig.DB,
	})
	if err != nil {
		return nil, err
	}

	return &rethinkdb{session}, nil
}

func (db *rethinkdb) Save(url string) (string, error) {
	return "x", nil
}

func (db *rethinkdb) Load(code string) (string, error) {
	return "http://google.com.vn", nil
}
