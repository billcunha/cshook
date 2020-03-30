package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	Player player `json:"player"`
}

type player struct {
	State state `json:"state"`
}

type state struct {
	Burning int `json:"burning"`
	Flashed int `json:"flashed"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

func receiveEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		return
	}

	dst := new(bytes.Buffer)
	json.Indent(dst, reqBody, "", "  ")
	fmt.Println(dst.String())

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

	if newEvent.Player.State.Burning > 0 {
		fmt.Println("pegou fogo!")
		requestBody, _ := json.Marshal(map[string]string{
			"url": "https://www.myinstants.com/media/sounds/pegando-fogo-bicho.mp3",
		})

		http.Post("http://localhost:9001/bot/play", "application/json", bytes.NewBuffer(requestBody))
	}

	if newEvent.Player.State.Flashed > 50 {
		fmt.Println("ceguin!")
		requestBody, _ := json.Marshal(map[string]string{
			"url": "https://www.myinstants.com/media/sounds/derruba.mp3",
		})

		http.Post("http://localhost:9001/bot/play", "application/json", bytes.NewBuffer(requestBody))
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/bot", receiveEvent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
