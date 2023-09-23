package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/lapayka/rsoi/BL"
	"github.com/lapayka/rsoi/DA"
)

type GateWay struct {
	db     *DA.DB
	logger *slog.Logger
}

func CreateGateWay() (GateWay, error) {
	gw := GateWay{}
	gw.logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	var err error
	gw.db, err = DA.New()

	if err != nil {
		return GateWay{}, err
	}

	return gw, err
}

func main() {
	router := mux.NewRouter()

	gw, _ := CreateGateWay()

	router.HandleFunc("/persons", gw.getPerson).Methods("Get")
	router.HandleFunc("/persons/{id}", gw.getPersonById).Methods("GET")
	router.HandleFunc("/persons", gw.createPerson).Methods("POST")
	router.HandleFunc("/persons/{id}", gw.updatePerson).Methods("PATCH")
	router.HandleFunc("/persons/{id}", gw.deletePerson).Methods("DELETE")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		gw.logger.Error("failed to run app", "error", err)
	}
}

func (gw *GateWay) getPerson(w http.ResponseWriter, r *http.Request) {
	persons, _ := gw.db.GetPersons()

	if persons == nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		writeSerializable(&persons, w)
	}
}

func (gw *GateWay) getPersonById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	person, _ := gw.db.GetPersonById(int(id))

	err_eq := BL.Person{}
	if person == err_eq {
		w.WriteHeader(http.StatusNotFound)
	} else {
		writeSerializable(&person, w)
	}
}

func (gw *GateWay) createPerson(w http.ResponseWriter, r *http.Request) {
	p := BL.Person{}
	readSerializable(r, &p)

	err := gw.db.CreatePerson(&p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Add("Location", fmt.Sprintf("persons/%d", p.ID))
	w.WriteHeader(http.StatusCreated)
}

func (gw *GateWay) updatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	p := BL.Person{}
	readSerializable(r, &p)

	_ = gw.db.UpdatePerson(int(id), p)
}

func (gw *GateWay) deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	err := gw.db.DeletePerson(int(id))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

func readSerializable(r *http.Request, item BL.Serializable) {
	buff, _ := io.ReadAll(r.Body)

	err := item.FromJSON(string(buff))

	if err != nil {
		slog.Warn("Wrong http request", "parse error", err)
	}
}

func writeSerializable(item BL.Serializable, w http.ResponseWriter) {
	w.Write([]byte(item.ToJSON()))
}
