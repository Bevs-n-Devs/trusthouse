package handler

import (
	"net/http"

	"github.com/Bevs-n-Devs/trusthouse/logs"
	"github.com/Bevs-n-Devs/trusthouse/tmpl"
)

func Home(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		logs.Logs(logErr, "Unable to load home page: "+err.Error())
		http.Error(w, "Unable to load home page: "+err.Error(), http.StatusInternalServerError)
	}
}
