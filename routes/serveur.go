package routes

import (
	"fmt"
	"net/http"
)

func InitServ() {
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	accueilRoutes()
	rechercherRoutes()
	setRoutes()
	ProposRoutes()
	CardsBySetRoutes()
	fmt.Println("Serveur démarré sur http://localhost:8080/accueil")
	http.ListenAndServe(":8080", nil)
}
