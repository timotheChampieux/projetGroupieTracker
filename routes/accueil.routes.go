package routes

import (
	"GroupieTracker/services"
	"GroupieTracker/templates"
	"fmt"
	"net/http"
	"strconv"
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
		var dataTemp services.Pokemon
		query := r.FormValue("query")
		minHp := r.FormValue("min-hp")
		fmt.Println(minHp, query)

		var err error
		data, err = services.RecherchePokemon(query)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		minHpInt, err := strconv.Atoi(minHp)
		if err == nil && minHp != "" {
			for _, item := range data.Data {
				hpItemInt, hpErr := strconv.Atoi(item.Hp)
				if hpErr != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				fmt.Println(hpItemInt, minHpInt)
				if hpItemInt >= minHpInt {

					dataTemp.Data = append(dataTemp.Data, item)
				}
			}
		} else {

			dataTemp = data
		}
		for _, item := range data.Data {
			fmt.Println(item)
		}
		templates.Temp.ExecuteTemplate(w, "rechercher", dataTemp)
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

		list, err := services.GetCardsBySet(id)
		data := services.CardPokemonFinal{name, list}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(data)
		templates.Temp.ExecuteTemplate(w, "cardset", data)
	})

}
