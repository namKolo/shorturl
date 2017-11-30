package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	storage "github.com/namKolo/shorturl/storage"
)

// ItemHandler take care the item routes
type ItemHandler struct {
	prefix  string
	storage storage.Storage
}

// NewItemHandler returns new ItemHandler
func NewItemHandler(prefix string, storage storage.Storage) *ItemHandler {
	return &ItemHandler{prefix, storage}
}

// Encode is used to encode URL
func (h ItemHandler) Encode(w io.Writer, r *http.Request) (interface{}, int, error) {
	if r.Method != http.MethodPost {
		return nil, http.StatusMethodNotAllowed, fmt.Errorf("Method %s not allowed", r.Method)
	}

	var input struct{ URL string }
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("Unable to decode JSON request body: %v", err)
	}

	c, err := h.storage.Save(input.URL)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("Could not save to db: %v", err)
	}

	return h.prefix + c, http.StatusCreated, nil
}

// Redirect is used to redirect
func (h ItemHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	code := r.URL.Path[len("/"):]

	url, err := h.storage.Load(code)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("URL not found"))
		return
	}
	http.Redirect(w, r, string(url), http.StatusMovedPermanently)
}
