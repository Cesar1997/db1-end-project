package main

import (
	"encoding/json"
	"net/http"

	"github.com/Cesar1997/db1-end-project/structures"
)

func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//Reading JSON Request
	var request structures.User

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		responseWithError(w, r, err)
		return
	}
	err = createUser(request)

	if err != nil {
		responseWithError(w, r, err)
		return
	}

	res := structures.Response
	res.Message = "usuario registrado con éxito"
	json.NewEncoder(w).Encode(res)

}

func loginUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var request structures.User
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		responseWithError(w, r, err)
		return
	}
	userLogged, err := loginUser(request)

	if err != nil {
		responseWithError(w, r, err)
		return
	}

	res := structures.Response

	if userLogged == (structures.User{}) {
		res.Result = false
		res.Message = "No existe el usuario"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Records = userLogged
	res.Message = "incio sesión exitosamente"
	json.NewEncoder(w).Encode(res)
}
