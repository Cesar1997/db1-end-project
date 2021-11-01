package main

import (
	"encoding/json"
	"net/http"

	"github.com/Cesar1997/db1-end-project/db"
	"github.com/Cesar1997/db1-end-project/structures"
)

func getAllCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	listCountries, err := db.GetAllCountries()
	if err != nil {
		responseWithError(w, r, err)
		return
	}
	res := structures.Response
	res.Message = "Paises consultados con éxito"
	res.Records = listCountries
	json.NewEncoder(w).Encode(res)

}
