package main

import (
	"encoding/json"
	"net/http"

	"github.com/Cesar1997/db1-end-project/db"
	"github.com/Cesar1997/db1-end-project/structures"
)

func storePurchaseOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var request structures.PurchaseOrderType
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		responseWithError(w, r, err)
		return
	}
	err = db.CreatePurchaseOrder(request)

	if err != nil {
		responseWithError(w, r, err)
		return
	}

	res := structures.Response
	res.Message = "Se creo exitosamente la orden de alquiler"
	json.NewEncoder(w).Encode(res)
}
