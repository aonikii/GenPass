package user

import (
	"GenPass/internal/sessions"
	"html/template"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request, tmpl *template.Template) {
	switch r.Method {
	case "GET":
		if sessions.CheckUserSession(w, r) {
			tmpl.ExecuteTemplate(w, "dashboard.html", nil)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}

}
