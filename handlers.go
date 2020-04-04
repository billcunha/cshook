package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		fmt.Fprintf(w, "Awesome fail")
		return
	}

	dst := new(bytes.Buffer)
	json.Indent(dst, reqBody, "", "  ")
	logger.Println(dst.String())

	if reqBody == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "You need to send some data")
		return
	}

	err = json.Unmarshal(reqBody, &newEvent)
	if err != nil {
		fmt.Fprintf(w, "JSON decode fail")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Disable all events from other players
	if newEvent.Provider.Steamid != newEvent.Player.Steamid {
		return
	}

	// Disable all events from other places
	if newEvent.Player.Activity != "playing" {
		return
	}

	CheckBurning(newEvent)
	CheckFlashed(newEvent)
	CheckDead(newEvent)
	CheckHeadShot(newEvent)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok")
}
