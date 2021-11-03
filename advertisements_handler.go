package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Cesar1997/db1-end-project/db"
	"github.com/Cesar1997/db1-end-project/structures"
	"github.com/gorilla/mux"
)

func getAdvertisementsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	listAdds, err := db.GetAllAdds(id)

	if err != nil {
		responseWithError(w, r, err)
		return
	}

	res := structures.Response
	res.Message = "Listado de anuncios obtenidos exitosamente"
	res.Records = listAdds
	json.NewEncoder(w).Encode(res)
}

func getAdvertisementHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	add, err := db.GetAdd(id)
	if err != nil {
		responseWithError(w, r, err)
		return
	}
	res := structures.Response
	res.Message = "Anuncio consultado con éxito"
	res.Records = add
	json.NewEncoder(w).Encode(res)
}

func storeAdvertisementHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var request structures.Adds
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		responseWithError(w, r, err)
		return
	}
	err = db.CreateAdd(request)
	if err != nil {
		responseWithError(w, r, err)
		return
	}

	res := structures.Response
	res.Message = "Anuncio creado exitosamente"
	json.NewEncoder(w).Encode(res)
}

func getCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	add, err := db.GetAllCategories()
	if err != nil {
		responseWithError(w, r, err)
		return
	}
	res := structures.Response
	res.Message = "Categorias consultadas exitosamente"
	res.Records = add
	json.NewEncoder(w).Encode(res)

}
