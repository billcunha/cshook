package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

// CheckBurning ...
func CheckBurning(event Event) {
	if event.Player.State.Burning == 0 {
		return
	}

	sound := getSound(config.Burning)
	sendRequest(sound)
}

// CheckFlashed ...
func CheckFlashed(event Event) {
	if event.Player.State.Flashed < 50 {
		return
	}

	sound := getSound(config.Flashed)
	sendRequest(sound)
}

// CheckDead ...
func CheckDead(event Event) {
	if event.Previously.Player.State.Health == 0 || event.Player.State.Health != 0 {
		return
	}

	sound := getSound(config.Dead)
	sendRequest(sound)
}

// CheckHeadShot ...
func CheckHeadShot(event Event) {
	if event.Previously.Player.State.RoundKillhs >= event.Player.State.RoundKillhs {
		return
	}

	for key, value := range event.Previously.Player.Weapons {
		if (value.AmmoClip + 1) == event.Player.Weapons[key].AmmoClip {
			sound := getSound(config.Headshot)
			sendRequest(sound)
			return
		}
	}
}

func getSound(sounds []string) string {
	return sounds[rand.Intn(len(sounds))]
}

func sendRequest(sound string) {
	fmt.Println(sound)
	requestBody, _ := json.Marshal(map[string]string{
		"url": sound,
	})

	http.Post(config.BotAddress, "application/json", bytes.NewBuffer(requestBody))
}
