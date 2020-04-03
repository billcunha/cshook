package main

// TomlConfig ...
type TomlConfig struct {
	Port       string
	BotAddress string `toml:"bot_address"`
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
	Player     Player     `json:"player"`
	Previously Previously `json:"previously"`
}

// Player ...
type Player struct {
	State State `json:"state"`
}

// Previously ...
type Previously struct {
	Player Player `json:"player"`
}

// State ...
type State struct {
	Burning int `json:"burning"`
	Flashed int `json:"flashed"`
	Health  int `json:"health"`
}
