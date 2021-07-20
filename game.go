package main

type GameParam struct {
	Moves     bool
	PgnInJson bool
	Tags      bool
	Clocks    bool
	Evals     bool
	Opening   bool
	Literate  bool
	Players   string
}

func NewGameParam() GameParam {
	return GameParam{
		Moves:     true,
		PgnInJson: false,
		Tags:      true,
		Clocks:    true,
		Evals:     true,
		Opening:   true,
		Literate:  false,
		Players:   "",
	}
}

type Game struct {
	ID         string `json:"id"`
	Rated      bool   `json:"rated"`
	Variant    string `json:"variant"`
	Speed      string `json:"speed"`
	Perf       string `json:"perf"`
	CreatedAt  int64  `json:"createdAt"`
	LastMoveAt int64  `json:"lastMoveAt"`
	Status     string `json:"status"`
	Players    struct {
		White struct {
			User struct {
				Name   string `json:"name"`
				Title  string `json:"title"`
				Patron bool   `json:"patron"`
				ID     string `json:"id"`
			} `json:"user"`
			Rating     int `json:"rating"`
			RatingDiff int `json:"ratingDiff"`
		} `json:"white"`
		Black struct {
			User struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"user"`
			Rating     int `json:"rating"`
			RatingDiff int `json:"ratingDiff"`
		} `json:"black"`
	} `json:"players"`
	Opening struct {
		Eco  string `json:"eco"`
		Name string `json:"name"`
		Ply  int    `json:"ply"`
	} `json:"opening"`
	Moves string `json:"moves"`
	Clock struct {
		Initial   int `json:"initial"`
		Increment int `json:"increment"`
		TotalTime int `json:"totalTime"`
	} `json:"clock"`
}
