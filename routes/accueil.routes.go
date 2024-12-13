package routes

import (
	"GroupieTracker/templates"
	"net/http"
)

func accueilRoutes() {
	// Route for the home page
	http.HandleFunc("/accueil", func(w http.ResponseWriter, r *http.Request) {
		templates.Temp.ExecuteTemplate(w, "accueil", nil)
	})
}
