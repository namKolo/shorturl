package rethinkdb

import (
	r "github.com/dancannon/gorethink"
	config "github.com/namKolo/shorturl/config"
	model "github.com/namKolo/shorturl/model"
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
	item := model.NewItem(url)
	res, err := r.Table("items").Insert(item).RunWrite(db.session)
	if err != nil {
		return "", err
	}
	id := res.GeneratedKeys[0]

	return id, nil
}

func (db *rethinkdb) Load(code string) (string, error) {
	res, err := r.Table("items").Get(code).Run(db.session)
	if err != nil {
		return "", err
	}
	var item model.Item
	err = res.One(&item)
	if err != nil {
		return "", err
	}

	return item.URL, nil
}
