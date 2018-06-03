package main

type DBApiResponse struct {
	Type          string    `json:"type"`
	Number        string    `json:"number"`
	Name          string    `json:"name"`
	Stations      []Station `json:"stations"`
	ValidUntil    int       `json:"validUntil"`
	Jobs          []string  `json:"jobs"`
	Tickermessage string    `json:"tickermessage"`
	Afz struct {
		Bikes      int `json:"bikes"`
		Passengers int `json:"passengers"`
		Capacity struct {
			Passengers int `json:"passengers"`
			Bikes      int `json:"bikes"`
		} `json:"capacity"`
	} `json:"afz"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Speed     float64     `json:"speed"`
	Timestamp int     `json:"timestamp"`
	Cards     []Card  `json:"cards"`
}

type Station struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	ShortTag  string  `json:"shortTag"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Track     int     `json:"track"`
	Approach  int     `json:"approach"`
	Departure int     `json:"departure"`
	Delay     int     `json:"delay"`
	Status    string  `json:"status"`
	Message   string  `json:"message"`
	Distances struct {
		Next float64 `json:"next"`
		Prev int     `json:"prev"`
	} `json:"distances"`
}

type Card struct {
	Sim      string `json:"sim"`
	Type     string `json:"type"`
	Signal   int    `json:"signal"`
	Provider string `json:"provider"`
}
