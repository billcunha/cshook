package main

// TomlConfig ...
type TomlConfig struct {
	Port       string
	BotAddress string
	Player     struct {
		State struct {
			Burning []string
			Flashed []string
			Dead    []string
		}
	}
}

// Event ... Root CS event
type Event struct {
	Player Player `json:"player"`
}

// Player ...
type Player struct {
	State State `json:"state"`
}

// State ...
type State struct {
	Burning int `json:"burning"`
	Flashed int `json:"flashed"`
	Health  int `json:"health"`
}
