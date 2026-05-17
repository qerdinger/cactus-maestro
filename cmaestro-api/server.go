package main

import (
	"cmaestro-api/internal/router"
	cmaestro_db "cmaestro-db"
	cregistry "cmastero-registry"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

/*
Infrastructure Requirements

|------------------- Infrastructure ------------------|
| Docker Registry | image=registry:3 | exposing=:5001 |
| Redis DB 	   	  | image=redis	     | exposing=:6379 |
|-----------------------------------------------------|

*/

func main() {
	r := router.NewRouter()

	// Middlewares Configuration
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/users", getUsers)

	// **************************************** SERVICES ****************************************
	// system
	// health
	// status

	// registry
	// List every image hosted on local docker registry

	// registry/create (update-in-place)
	// registry/list
	// registry/delete
	// registry/_/function_path

	// insight/
	// insight/_/function_path
	// ******************************************************************************************
	// All these previous methods could be authenticated and authorised (feature available later)
	// ******************************************************************************************

	//r.ListenAndServe(":8080")

	/*cactuizedFunctions := cingest.Ingest(`from cactuskit import ApiMethod, ApiProtocol, HttpStatus, cactuize

	@cactuize()
	def simple_entrypoint():
	    return f"Hello World from {simple_entrypoint}"`)

		//fmt.Println(cactuizedFunctions)*/

	registry := cregistry.New("http://localhost:5001", nil, nil)
	catalog, _ := registry.GetCatalog()
	fmt.Println("Catalog : ", catalog)
	//dig, err := registry.GetDigest("registry", "latest")
	//fmt.Println(dig, err)
	//status, err := registry.RemoveTag("registry", "latest")
	//fmt.Println("Has been deleted?", status, err)

	db, _ := cmaestro_db.New(cmaestro_db.Config{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := db.Set("hello", "world", -1)
	if err != nil {
		return
	}

	k := "hello"
	v, err := db.Get(k)
	fmt.Printf("Value as key='{%s}' %s | err={%v}", k, v, err)
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	users := []string{"admin"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
