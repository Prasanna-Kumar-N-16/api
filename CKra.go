package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Comp struct {
	Cname string          `json:"cname"`
	Kra   map[string]Kras `json:"kra"`
}
type Kras struct {
	Title string            `json:"title"`
	Kpi   map[string]string `json:"kpi"`
}

var comp []Comp

func Company(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(comp)
}

func Getkra(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for _, y := range comp {
		for kr, _ := range y.Kra {
			if kr == vars["kra"] {
				json.NewEncoder(w).Encode(y.Kra[kr])
				return
			}
		}
	}
	json.NewEncoder(w).Encode(&Comp{})
}
func NewKra(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Comp
	json.NewDecoder(r.Body).Decode(&book)

	for _, y := range comp {
		for a, b := range book.Kra {
			y.Kra[a] = b
		}
	}
	json.NewEncoder(w).Encode(book.Kra)
}
func UpdateKra(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var book Comp
	json.NewDecoder(r.Body).Decode(&book)
	for _, y := range comp {
		for kr, _ := range y.Kra {
			for a, b := range book.Kra {
				if kr == vars["kra"] {
					y.Kra[a] = b
				}
			}
		}
	}
	json.NewEncoder(w).Encode(book.Kra)
}
func DeleteKra(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for _, y := range comp {
		for kr, _ := range y.Kra {
			if kr == vars["kra"] {
				json.NewEncoder(w).Encode(y.Kra[kr])
				delete(y.Kra, kr)
			}
		}
	}

}
func main() {
	r := mux.NewRouter()
	cm := new(Comp)
	cm.Cname = "Param"
	cm.Kra = make(map[string]Kras)
	e, f := cm.Kra["K1"]
	if !f {
		e = Kras{}
	}
	e.Title = "Title Description"
	e.Kpi = make(map[string]string)
	cm.Kra["K1"] = e
	cm.Kra["K1"].Kpi["KPI1"] = "Metrics One"
	cm.Kra["K1"].Kpi["KPI2"] = "Metrics Two"
	comp = append(comp, Comp{Cname: cm.Cname, Kra: cm.Kra})
	r.HandleFunc("/cmp", Company).Methods("GET")
	r.HandleFunc("/kra/{kra}", Getkra).Methods("GET")
	r.HandleFunc("/newkra/Comp", NewKra).Methods("POST")
	r.HandleFunc("/updkra/{kra}", UpdateKra).Methods("PUT")
	r.HandleFunc("/delete/{kra}", DeleteKra).Methods("DELETE")
	http.Handle("/", r)
	http.ListenAndServe(":1000", nil)
}
