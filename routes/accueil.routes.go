package routes

import (
	"GroupieTracker/services"
	"GroupieTracker/templates"
	"fmt"
	"math"
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
		maxHp := r.FormValue("max-hp")
		degMin := r.FormValue("deg-min")
		fmt.Println(minHp, maxHp, query)

		if query != "" {
			var err error
			data, err = services.RecherchePokemon(query)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			minHpInt, errMin := strconv.Atoi(minHp)
			maxHpInt, errMax := strconv.Atoi(maxHp)
			degMinInt, errDeg := strconv.Atoi(degMin)

			for _, item := range data.Data {

				hpItemInt, _ := strconv.Atoi(item.Hp)
				var dgItemInt int
				var errDegItem error = fmt.Errorf("Not")
				if len(item.Attacks) > 0 {
					dgItemInt, errDegItem = strconv.Atoi(item.Attacks[0].Damage)
					if errDegItem != nil {
						fmt.Println("2")
						fmt.Println(errDegItem)
					}
				}

				filtreHpMin := minHp == "" || (errMin == nil && minHpInt <= hpItemInt)
				filtreHpMax := maxHp == "" || (errMax == nil && maxHpInt >= hpItemInt)
				filtreDegMin := degMin == "" || (errDeg == nil && degMinInt <= dgItemInt)
				fmt.Println(degMinInt, dgItemInt, filtreDegMin)

				if filtreHpMin && filtreHpMax && filtreDegMin {
					dataTemp.Data = append(dataTemp.Data, item)
				}
			}

		}
		templates.Temp.ExecuteTemplate(w, "rechercher", dataTemp)
	})
}

func setRoutes() {
	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		var data services.SetPage
		sets, err := services.GetSet()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		taille := len(sets.Data)
		data.NbrPage = int(math.Ceil(float64(taille) / float64(8))) // division arrondie a l'entier supp

		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 1 {
			page = 1
		} else if page > data.NbrPage {
			page = data.NbrPage
		}

		data.PageAct = page

		data.PrevPage = page - 1
		data.NextPage = page + 1
		data.ShowPrev = page > 1
		data.ShowNext = page < data.NbrPage

		index := data.PageAct*8 - 8

		for i := 0; i < 8; i++ {

			data.Set.Data = append(data.Set.Data, sets.Data[index])
			index++
		}

		templates.Temp.ExecuteTemplate(w, "set", data)
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

func ProposRoutes() {
	http.HandleFunc("/propos", func(w http.ResponseWriter, r *http.Request) {
		templates.Temp.ExecuteTemplate(w, "propos", nil)
	})
}
