package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/rocket4ce/go/service"
)

type Custumer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustumerHandlers struct {
	service service.CustomerService
}

func (ch *CustumerHandlers) getAllCustumers(w http.ResponseWriter, r *http.Request) {
	// custumers := []Custumer{
	// 	{Name: "Dinko", City: "Coquimbo", ZipCode: "123123"},
	// 	{Name: "Dinko2", City: "Coquimbo", ZipCode: "234563"},
	// }

	custumers, _ := ch.service.GetAllCustumers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(custumers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(custumers)

	}
}

// func getCustumer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["curtumer_id"])
// }

// func createCustumer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "post request")
// }
