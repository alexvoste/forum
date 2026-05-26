package handler

import (
	"context"
	"database/sql"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		username := strings.TrimSpace(r.FormValue("username"))
		email := strings.TrimSpace(r.FormValue("email"))
		pubkey := strings.TrimSpace(r.FormValue("pubkey"))
		password := strings.TrimSpace(r.FormValue("password"))

		if username == "" || email == "" {
			http.Error(w, "username and email required", http.StatusBadRequest)
			return
		}

		if !strings.Contains(email, "@") {
			http.Error(w, "invalid email", http.StatusBadRequest)
			return
		}

		if pubkey == "" && password == "" {
			http.Error(w, "pubkey or password required", http.StatusBadRequest)
			return
		}

		var passwordHash string
		if password != "" {
			hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, "Internal error", http.StatusInternalServerError)
				return
			}
			passwordHash = string(hash)
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		query := `INSERT INTO users (username, email, ssh_pubkey, password_hash) VALUES (?, ?, ?, ?)`
		_, err := db.ExecContext(ctx, query, username, email, pubkey, passwordHash)
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				http.Error(w, "username or email exists", http.StatusConflict)
				return
			}
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
