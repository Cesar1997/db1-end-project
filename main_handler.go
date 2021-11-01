package main

import (
	"encoding/json"
	"net/http"

	"github.com/Cesar1997/db1-end-project/structures"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	response := structures.RESTResponse{
		Result:  true,
		Message: "The api server is running ...",
	}

	json.NewEncoder(w).Encode(response)
}
