package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Cesar1997/db1-end-project/db"
	"github.com/Cesar1997/db1-end-project/structures"
	"github.com/gorilla/mux"
)

func getAllProductsFilteredByuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["profileID"])

	products, err := db.GetAllProductsFilteredByUser(id)
	if err != nil {
		responseWithError(w, r, err)
		return
	}
	res := structures.Response
	res.Message = "productos consultados con éxito"
	res.Records = products
	json.NewEncoder(w).Encode(res)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	//GetProduct
	w.Header().Set("Access-Control-Allow-Origin", "*")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["productID"])

	products, err := db.GetProduct(id)
	if err != nil {
		responseWithError(w, r, err)
		return
	}
	res := structures.Response
	res.Message = "producto consultado con éxito"
	res.Records = products
	json.NewEncoder(w).Encode(res)
}

func storeProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var request structures.ProductType
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		responseWithError(w, r, err)
		return
	}
	err = db.CreateProduct(request)

	if err != nil {
		responseWithError(w, r, err)
		return
	}

	res := structures.Response

	res.Message = "Se creo exitosamente el producto"
	json.NewEncoder(w).Encode(res)
}

func getAllSizes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	listSizes, err := db.GetAllSizes()
	if err != nil {
		responseWithError(w, r, err)
		return
	}
	res := structures.Response
	res.Message = "productos consultados con éxito"
	res.Records = listSizes
	json.NewEncoder(w).Encode(res)

}

func getAllColors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	listColors, err := db.GetAllColors()
	if err != nil {
		responseWithError(w, r, err)
		return
	}
	res := structures.Response
	res.Message = "productos consultados con éxito"
	res.Records = listColors
	json.NewEncoder(w).Encode(res)

}

func getAllTypeArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	listTypeArticles, err := db.GetAllTypeArticle()
	if err != nil {
		responseWithError(w, r, err)
		return
	}
	res := structures.Response
	res.Message = "productos consultados con éxito"
	res.Records = listTypeArticles
	json.NewEncoder(w).Encode(res)
}
