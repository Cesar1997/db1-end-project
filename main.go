package main

import (
	"log"
	"net/http"

	"github.com/Cesar1997/db1-end-project/db"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

//Root server
func main() {

	/* Por el momento no se le puso ningún middleware de rutas
	ni una sesión o jwt . Es simple :c
	*/

	err := db.ConnectToDB()
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler)

	//Catalogs product
	router.HandleFunc("/color/list", getAllColors).Methods(http.MethodGet)
	router.HandleFunc("/size/list", getAllSizes).Methods(http.MethodGet)
	router.HandleFunc("/typeProduct/list", getAllTypeArticles).Methods(http.MethodGet)
	router.HandleFunc("/country/list", getAllCountries).Methods(http.MethodGet)
	//User interaction
	router.HandleFunc("/user/register", registerUserHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/user/login", loginUserHandler).Methods(http.MethodPost, http.MethodOptions)

	router.HandleFunc("/user/advertisements/create", storeAdvertisementHandler).Methods(http.MethodPost)
	router.HandleFunc("/product/advertisements/advertisements/list/{id}", getAdvertisementsHandler).Methods(http.MethodGet)
	router.HandleFunc("/product/advertisements/{id}", getAdvertisementHandler).Methods(http.MethodGet)
	router.HandleFunc("/advertisements/categories", getCategoryHandler).Methods(http.MethodGet)

	router.HandleFunc("/product", storeProduct).Methods(http.MethodPost)

	router.HandleFunc("/user/products/{profileID}", getAllProductsFilteredByuser).Methods(http.MethodGet)
	router.HandleFunc("/user/product/{productID}", getProduct).Methods(http.MethodGet)

	router.HandleFunc("/purchaseOrder", storePurchaseOrder).Methods(http.MethodPost)

	router.HandleFunc("/purchaseOrder/report/{id}", getReportPurchaseOrder).Methods(http.MethodGet)
	//Run server in port 8081
	log.Fatal(http.ListenAndServe(":8081", router))
}
