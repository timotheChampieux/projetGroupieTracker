package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Pokemon struct {
	Data []struct {
		ID       string   `json:"id"`
		Name     string   `json:"name"`
		Subtypes []string `json:"subtypes"`
		Hp       string   `json:"hp"`
		Types    []string `json:"types"`
		Rarity   string   `json:"rarity"`
		Images   struct {
			Small string `json:"small"`
		} `json:"images"`
	} `json:"data"`
}

var _token string = "0f0a0dd2-300f-4af3-9c5b-4bdb8270e451"

func RecherchePokemon(pokemon string) (Pokemon, error) {

	urlApi := "https://api.pokemontcg.io/v2/cards?q=name:" + pokemon

	httpClient := http.Client{
		Timeout: time.Second * 5,
	}

	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		return Pokemon{}, fmt.Errorf("Requete - Erreur lors de l'initialisation de la requête : %v", errReq)
	}

	req.Header.Set("X-Api-Key", _token)

	res, errRes := httpClient.Do(req)
	if errRes != nil {
		return Pokemon{}, fmt.Errorf("Requete - Erreur lors de l'envoi de la requête : %v", errRes)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("Requete - Erreur code : %v message : %v", res.StatusCode, res.Status)
	}

	var data Pokemon
	errDecode := json.NewDecoder(res.Body).Decode(&data)
	if errDecode != nil {
		return Pokemon{}, fmt.Errorf("Requete - Erreur lors de la lecture du JSON : %v", errDecode)
	}
	return data, nil
}
