package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	config "github.com/namKolo/shorturl/config"
	storage "github.com/namKolo/shorturl/storage"
	rethinkdb "github.com/namKolo/shorturl/storage/rethinkdb"
	handler "github.com/namkolo/shorturl/handler"
	mdw "github.com/namkolo/shorturl/middleware"
)

type App struct {
	Router  *mux.Router
	Storage storage.Storage
	Config  *config.App
}

func NewApp(config *config.App) (*App, error) {
	var app App
	db, err := rethinkdb.New(&config.RethinkDB)

	if err != nil {
		return nil, err
	}

	app.Storage = db
	app.Router = mux.NewRouter()
	app.Config = config
	app.setRouters()
	return &app, nil
}

func (a *App) setRouters() {
	itemHandlers := handler.NewItemHandler(a.Config.Options.Prefix, a.Storage)
	a.Post("/encode", mdw.ResponseHandler(itemHandlers.Encode))
	a.Get("/{id}", itemHandlers.Redirect)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":"+a.Config.Server.Port, a.Router))
}
