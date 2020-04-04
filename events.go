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

	sound := config.Burning[rand.Intn(len(config.Burning))]
	fmt.Println(sound)
	requestBody, _ := json.Marshal(map[string]string{
		"url": sound,
	})

	http.Post(config.BotAddress, "application/json", bytes.NewBuffer(requestBody))
}

// CheckFlashed ...
func CheckFlashed(event Event) {
	if event.Player.State.Flashed < 50 {
		return
	}

	sound := config.Flashed[rand.Intn(len(config.Flashed))]
	fmt.Println(sound)
	requestBody, _ := json.Marshal(map[string]string{
		"url": sound,
	})

	http.Post(config.BotAddress, "application/json", bytes.NewBuffer(requestBody))
}

// CheckDead ...
func CheckDead(event Event) {
	if event.Previously.Player.State.Health == 0 || event.Player.State.Health != 0 {
		return
	}

	sound := config.Dead[rand.Intn(len(config.Dead))]
	fmt.Println(sound)
	requestBody, _ := json.Marshal(map[string]string{
		"url": sound,
	})

	http.Post(config.BotAddress, "application/json", bytes.NewBuffer(requestBody))
}

// CheckHeadShot ...
func CheckHeadShot(event Event) {
	if event.Previously.Player.State.RoundKillhs >= event.Player.State.RoundKillhs {
		return
	}

	for key, value := range event.Previously.Player.Weapons {
		if (value.AmmoClip + 1) == event.Player.Weapons[key].AmmoClip {
			sound := config.Headshot[rand.Intn(len(config.Headshot))]
			fmt.Println(sound)
			requestBody, _ := json.Marshal(map[string]string{
				"url": sound,
			})

			http.Post(config.BotAddress, "application/json", bytes.NewBuffer(requestBody))

			return
		}
	}
}
