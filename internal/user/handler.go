package user

import (
	"GenPass/internal/models"
	"GenPass/internal/sessions"
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	switch r.Method {
	case "GET":
		if sessions.CheckUserSession(w, r) {
			sp := getPasswords(sessions.GetUserIdFromSession(w, r))
			tmpl.ExecuteTemplate(w, "dashboard.html", map[string][]models.Password{"Passwords": sp})
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}

}

func Generate(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	switch r.Method {
	case "GET":
		if !sessions.CheckUserSession(w, r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			http.Error(w, "сюда нельзя", 404)
		}
	case "POST":
		site := r.FormValue("site")
		length := r.FormValue("length")
		specialSymbol := r.FormValue("include_special")

		newPasswordCreation(site, length, specialSymbol, sessions.GetUserIdFromSession(w, r))

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}
