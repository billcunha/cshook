package main

// TomlConfig ...
type TomlConfig struct {
	Port       string
	BotAddress string `toml:"bot_address"`
	Burning    []string
	Flashed    []string
	Dead       []string
	Headshot   []string
}

// Event ... Root CS event
type Event struct {
	Player     Player     `json:"player"`
	Previously Previously `json:"previously"`
	Provider   struct {
		Steamid string
	}
}

// Player ...
type Player struct {
	Steamid  string
	State    State             `json:"state"`
	Activity string            `json:"activity"`
	Weapons  map[string]Weapon `json:"weapons"`
}

// Previously ...
type Previously struct {
	Player Player `json:"player"`
}

// State ...
type State struct {
	Burning     int `json:"burning"`
	Flashed     int `json:"flashed"`
	Health      int `json:"health"`
	RoundKillhs int `json:"round_killhs"`
}

// Weapon ...
type Weapon struct {
	AmmoClip int `json:"ammo_clip"`
}
