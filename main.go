package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

func receiveEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		return
	}

	if reqBody == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "You need to send some data...")
		return
	}

	err = json.Unmarshal(reqBody, &newEvent)
	if err != nil {
		fmt.Fprintf(w, "JSON decode fail")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if newEvent.ID != "" {
		fmt.Println(newEvent)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}

func main() {
	// initEvents()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/bot", receiveEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
