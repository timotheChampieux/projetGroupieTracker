package routes

import (
	"GroupieTracker/services"
	"GroupieTracker/templates"
	"net/http"
)

var (
	pokemon services.Pokemon
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

func rechercherRoutes() {

	http.HandleFunc("/rechercher", func(w http.ResponseWriter, r *http.Request) {
		var data services.Pokemon
		query := r.FormValue("query")
		if query != "" {
			var err error
			data, err = services.RecherchePokemon(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		templates.Temp.ExecuteTemplate(w, "rechercher", data)
	})
}
