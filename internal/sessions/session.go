package sessions

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

const (
	sessionMaxAge   = 7 * 24 * 60 * 60
	sessionSecure   = true
	sessionHttpOnly = true
)

var store *sessions.CookieStore

func SessionsInit() {
	key := os.Getenv("SESSION_KEY")
	if key == "" {
		log.Fatal("Не удалось достать ключ сессии")
	}

	store = sessions.NewCookieStore([]byte(key))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   sessionMaxAge,
		HttpOnly: sessionHttpOnly,
		Secure:   sessionSecure,
		SameSite: http.SameSiteLaxMode,
	}

}

func AddUserToSession(w http.ResponseWriter, r *http.Request, userId int) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Ошибка сессии", http.StatusInternalServerError)
		return
	}
	session.Values["user_id"] = userId

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Ошибка при сохранении сессии", http.StatusInternalServerError)
		return
	}
}

func CheckUserSession(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Ошибка сессии", http.StatusInternalServerError)
		return false
	}
	_, ok := session.Values["user_id"].(int)
	return ok
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Ошибка сессии при выходе", http.StatusInternalServerError)
		return
	}

	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Ошибка при сохранении сессии", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetUserIdFromSession(w http.ResponseWriter, r *http.Request) int {
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, "Ошибка сессии", http.StatusInternalServerError)
	}
	userId, _ := session.Values["user_id"].(int)
	return userId
}
