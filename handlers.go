package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

func receiveEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		return
	}

	// dst := new(bytes.Buffer)
	// json.Indent(dst, reqBody, "", "  ")
	// fmt.Println(dst.String())

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
		sound := config.Player.State.Burning[rand.Intn(len(config.Player.State.Burning))]
		fmt.Println(sound)
		requestBody, _ := json.Marshal(map[string]string{
			"url": sound,
		})

		http.Post(config.BotAddress, "application/json", bytes.NewBuffer(requestBody))
	}

	if newEvent.Player.State.Flashed > 50 {
		sound := config.Player.State.Flashed[rand.Intn(len(config.Player.State.Flashed))]
		fmt.Println(sound)
		requestBody, _ := json.Marshal(map[string]string{
			"url": sound,
		})

		http.Post(config.BotAddress, "application/json", bytes.NewBuffer(requestBody))
	}

	if newEvent.Player.State.Health == 0 {
		sound := config.Player.State.Dead[rand.Intn(len(config.Player.State.Dead))]
		fmt.Println(sound)
		requestBody, _ := json.Marshal(map[string]string{
			"url": sound,
		})

		http.Post(config.BotAddress, "application/json", bytes.NewBuffer(requestBody))
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}
