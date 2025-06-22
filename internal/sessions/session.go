package sessions

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func SessionsInit() {
	key := os.Getenv("SESSION_KEY")
	if key == "" {
		log.Fatal("Не удалось достать ключ сессии")
	}

	Store = sessions.NewCookieStore([]byte(key))
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
		Secure:   true,
		SameSite: 1,
	}

}

func AddUserToSession(w http.ResponseWriter, r *http.Request, userId int) {
	session, err := Store.Get(r, "session")
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
	session, err := Store.Get(r, "session")
	if err != nil {
		http.Error(w, "Ошибка сессии", http.StatusInternalServerError)
		return false
	}
	_, ok := session.Values["user_id"].(int)
	return ok
}
