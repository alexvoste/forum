package handler

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh"
)

func Login(db *sql.DB) http.HandlerFunc {
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

		log.Println("-> POST /login", username, email)

		if username == "" && email == "" {
			http.Error(w, "username or email required", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		var storedPubkey, storedHash string
		var query string
		var args []interface{}

		if username != "" {
			query = `SELECT ssh_pubkey, password_hash FROM users WHERE username = ?`
			args = []interface{}{username}
		} else {
			query = `SELECT ssh_pubkey, password_hash FROM users WHERE email = ?`
			args = []interface{}{email}
		}

		err := db.QueryRowContext(ctx, query, args...).Scan(&storedPubkey, &storedHash)
		if err == sql.ErrNoRows {
			http.Error(w, "user not found", http.StatusUnauthorized)
			return
		}
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		if pubkey != "" && storedPubkey != "" {
			if validatePubkey(pubkey, storedPubkey) {
				setSessionCookie(w)
				log.Println("Auth OK, rederection to /")
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			http.Error(w, "pubkey mismatch", http.StatusUnauthorized)
			return
		}

		if password != "" && storedHash != "" {
			if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err == nil {
				setSessionCookie(w)
				log.Println("Auth OK, rederection to /")
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
	}
}

func setSessionCookie(w http.ResponseWriter) {
	b := make([]byte, 32)
	rand.Read(b)
	token := hex.EncodeToString(b)

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "example_name",
		Value:    "example_value",
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func validatePubkey(provided, stored string) bool {
	pubKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(provided))
	if err != nil {
		return false
	}
	storedKey, _, _, _, err := ssh.ParseAuthorizedKey([]byte(stored))
	if err != nil {
		return false
	}
	return pubKey.Type() == storedKey.Type() &&
		string(pubKey.Marshal()) == string(storedKey.Marshal())
}
