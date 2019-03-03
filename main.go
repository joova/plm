package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	hdl "logika/plm/handlers"
)

func main() {
	// defer database.Disconnect()

	router := mux.NewRouter()
	//uom route
	router.HandleFunc("/api/plm/uoms", hdl.GetAllUOMEndpoint).Methods("GET")
	router.HandleFunc("/api/plm/uom/{code}", hdl.GetUOMEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/plm/uom/{code}", hdl.DeleteUOMEndpoint).Methods("DELETE")
	router.HandleFunc("/api/plm/uoms/{offset}/{limit}", hdl.GetPagingUOMEndpoint).Methods("GET")
	router.HandleFunc("/api/plm/uoms/{offset}/{limit}/{text}", hdl.SearchUOMEndpoint).Methods("GET")
	router.HandleFunc("/api/plm/uom", hdl.CreateUOMEndpoint).Methods("POST")
	router.HandleFunc("/api/plm/uom/{code}", hdl.UpdateUOMEndpoint).Methods("PUT")

	//category route
	router.HandleFunc("/api/plm/categories", hdl.GetAllCategoryEndpoint).Methods("GET")
	router.HandleFunc("/api/plm/category/{code}", hdl.GetCategoryEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/plm/category/{code}", hdl.DeleteCategoryEndpoint).Methods("DELETE")
	router.HandleFunc("/api/plm/categories/{offset}/{limit}", hdl.GetPagingCategoryEndpoint).Methods("GET")
	router.HandleFunc("/api/plm/categories/{offset}/{limit}/{text}", hdl.SearchCategoryEndpoint).Methods("GET")
	router.HandleFunc("/api/plm/category", hdl.CreateCategoryEndpoint).Methods("POST")
	router.HandleFunc("/api/plm/category/{code}", hdl.UpdateCategoryEndpoint).Methods("PUT")

	//org route
	router.HandleFunc("/api/plm/types", hdl.GetAllTypeEndpoint).Methods("GET")
	router.HandleFunc("/api/plm/type/{code}", hdl.CreateTypeEndpoint).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/plm/type/{code}", hdl.DeleteTypeEndpoint).Methods("DELETE")
	router.HandleFunc("/api/plm/types/{offset}/{limit}", hdl.GetPagingTypeEndpoint).Methods("GET")
	router.HandleFunc("/api/plm/types/{offset}/{limit}/{text}", hdl.SearchTypeEndpoint).Methods("GET")
	router.HandleFunc("/api/plm/type", hdl.CreateTypeEndpoint).Methods("POST")
	router.HandleFunc("/api/plm/type/{code}", hdl.UpdateTypeEndpoint).Methods("PUT")

	fmt.Println("Starting server on port 8001...")
	// log.Fatal(http.ListenAndServe(":8001", router))

	corsAllowedOriginsObj := handlers.AllowedOrigins([]string{"*"})
	corsAllowedHeadersObj := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	corsExposedHeadersObj := handlers.ExposedHeaders([]string{"Pagination-Count", "Pagination-Limit", "Pagination-Page"})
	corsAllowedMethodsObj := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":8001", handlers.CORS(corsAllowedOriginsObj, corsAllowedHeadersObj, corsExposedHeadersObj, corsAllowedMethodsObj)(router)))
}
