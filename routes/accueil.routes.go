package routes

import (
	"GroupieTracker/services"
	"GroupieTracker/templates"
	"fmt"
	"net/http"
)

var (
	pokemon services.Pokemon
)

func accueilRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

func setRoutes() {
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		sets, err := services.GetSet()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		templates.Temp.ExecuteTemplate(w, "set", sets)
	})
}

func CardsBySetRoutes() {
	http.HandleFunc("/set/cards", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		name := r.FormValue("name")
		fmt.Println(id, name)
		list, err := services.GetCardsBySet(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(list)
		templates.Temp.ExecuteTemplate(w, "cardset", list)
	})

}
