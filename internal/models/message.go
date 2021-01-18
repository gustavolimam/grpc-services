package models

type Message struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float32 `json:"coordinates"`
	Province    string    `json:"province"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

type MessageDB struct {
	ID          string `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Alias       string `json:"alias"`
	Regions     string `json:"regions"`
	Coordinates string `json:"coordinates"`
	Province    string `json:"province"`
	Timezone    string `json:"timezone"`
	Unlocs      string `json:"unlocs"`
	Code        string `json:"code"`
}
