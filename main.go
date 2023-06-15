package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/fernandomoranarita/Gobackend/db"
	"github.com/fernandomoranarita/Gobackend/handlers"
)

func main() {
	// activando las variables de entorno
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// llamando la base de datos
	db.DbConnection()

	// rutas
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	fmt.Print("\nServer Running on http://127.0.0.1:8000\n")
	http.ListenAndServe(":8000", router)
}
