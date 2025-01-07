package routes

import (
	"GroupieTracker/templates"
	"net/http"
)

func accueilRoutes() {
	// Route for the home page
	http.HandleFunc("/accueil", func(w http.ResponseWriter, r *http.Request) {
		/*listPokemon, err := services.RecherchePokemon("raichu")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}*/
		templates.Temp.ExecuteTemplate(w, "accueil", nil)
	})

}
