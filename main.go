package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/fernandomoranarita/Gobackend/db"
	"github.com/fernandomoranarita/Gobackend/handlers"
	"github.com/fernandomoranarita/Gobackend/helpers"
)

func main() {
	http.FileServer(http.Dir("templates"))
	// activando las variables de entorno
	err := godotenv.Load(".env")
	helpers.PanicErrorHandler(err) // <- manejador de errores para una mejor legibilidad del codigo

	// llamando la base de datos
	db.DbConnection()

	// rutas
	router := mux.NewRouter()

	// activando permisos cors
	var (
		ALLOWED_HOSTS   []string = strings.Split(os.Getenv("ALLOWED_HOSTS"), ",")
		ALLOWED_METHODS []string = strings.Split(os.Getenv("ALLOWED_METHODS"), ",")
	)

	c := cors.New(cors.Options{
		AllowedOrigins: ALLOWED_HOSTS, // <- strings.Split devuelve un slice de strings automaticamente es decir devuelve esto []string{devices}
		AllowedMethods: ALLOWED_METHODS,
	})

	// envolviendo el enrutador en el middlework
	routerMiddleware := c.Handler(router)

	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	fmt.Print("\nServer Running on http://127.0.0.1:8000\n")
	http.ListenAndServe(":8000", routerMiddleware)
}
