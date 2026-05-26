package routes

import (
	"database/sql"
	"net/http"

	"forum/internal/handler"
	"forum/internal/render"
)

var (
	main     = render.Handler
	pub      = render.PubHandler
	about    = render.AboutHandler
	register = render.RegisterHandler
	login    = render.LoginHandler
)

func InitRoutes(db *sql.DB) {
	http.HandleFunc("/", render.Handler)
	http.HandleFunc("/pub", render.PubHandler)
	http.HandleFunc("/about", render.AboutHandler)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			render.RegisterHandler(w, r)
		} else if r.Method == http.MethodPost {
			handler.Register(db)(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			render.LoginHandler(w, r)
		} else if r.Method == http.MethodPost {
			handler.Login(db)(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusInternalServerError)
		}
	})
}
