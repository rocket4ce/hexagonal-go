package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rocket4ce/go/domain"
	"github.com/rocket4ce/go/service"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	// wiring
	ch := CustumerHandlers{service.NewCustumerService(domain.NewCustumerRepositoryStub())}

	// rutas
	router.HandleFunc("/custumers", ch.getAllCustumers).Methods(http.MethodGet)
	// server
	log.Fatal(http.ListenAndServe(":8080", router))
}
