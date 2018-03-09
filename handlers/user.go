package handlers

import (
	"strconv"
	"strings"
	"github.com/rthornton128/test/services"
	"github.com/rthornton128/test/storage"
	"net/http"
)

// NewUserHandler returns a handler for the user service
func NewUserHandler() http.Handler {
	return &userHandler{}
}

type userHandler struct {}

func (u *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var s string
		var err error
		path := strings.Split(r.URL.Path, "/")
		userID := path[len(path)-1]
		if userID == "" {
			s, err = services.NewUserService(storage.NewUserStorage("db")).GetAll()
		}
		if len(path) > 2 && userID != "" {
			var id int64
			id, err = strconv.ParseInt(path[len(path)-1], 0, 64)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				s, err = services.NewUserService(storage.NewUserStorage("db")).Get(id)
			}
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Write([]byte(s))
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed),
			http.StatusMethodNotAllowed)
	}
}
