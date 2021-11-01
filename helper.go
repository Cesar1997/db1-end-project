package main

import (
	"encoding/json"
	"net/http"

	"github.com/Cesar1997/db1-end-project/structures"
)

func responseWithError(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(500)
	response := structures.RESTResponse{
		Result:  false,
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(response)
}
