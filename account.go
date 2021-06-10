package main

type Account struct {
	ID             string      `json:"id"`
	Username       string      `json:"username"`
	Online         bool        `json:"online"`
	Perfs          Performance `json:"perfs"`
	CreatedAt      int64       `json:"createdAt"`
	Disabled       bool        `json:"disabled"`
	TosViolation   bool        `json:"tosViolation"`
	Profile        Profile     `json:"profile"`
	SeenAt         int64       `json:"seenAt"`
	Patron         bool        `json:"patron"`
	PlayTime       PlayTime    `json:"playTime"`
	Language       string      `json:"language"`
	Title          string      `json:"title"`
	URL            string      `json:"url"`
	Playing        string      `json:"playing"`
	NbFollowing    int         `json:"nbFollowing"`
	NbFollowers    int         `json:"nbFollowers"`
	CompletionRate int         `json:"completionRate"`
	Count          Count       `json:"count"`
	Streaming      bool        `json:"streaming"`
	Followable     bool        `json:"followable"`
	Following      bool        `json:"following"`
	Blocking       bool        `json:"blocking"`
	FollowsYou     bool        `json:"followsYou"`
	Email          string      `json:",omitempty"`
}

type Email struct {
	Email string `json:"email"`
}

type Performance struct {
	Chess960 struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"chess960"`
	Atomic struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"atomic"`
	RacingKings struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"racingKings"`
	UltraBullet struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"ultraBullet"`
	Blitz struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"blitz"`
	KingOfTheHill struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"kingOfTheHill"`
	Bullet struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"bullet"`
	Correspondence struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"correspondence"`
	Horde struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"horde"`
	Puzzle struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"puzzle"`
	Classical struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"classical"`
	Rapid struct {
		Games  int  `json:"games"`
		Rating int  `json:"rating"`
		Rd     int  `json:"rd"`
		Prog   int  `json:"prog"`
		Prov   bool `json:"prov"`
	} `json:"rapid"`
	Storm struct {
		Runs  int `json:"runs"`
		Score int `json:"score"`
	} `json:"storm"`
}

type Profile struct {
	Country    string `json:"country"`
	Location   string `json:"location"`
	Bio        string `json:"bio"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	FideRating int    `json:"fideRating"`
	UscfRating int    `json:"uscfRating"`
	EcfRating  int    `json:"ecfRating"`
	Links      string `json:"links"`
}

type PlayTime struct {
	Total int `json:"total"`
	Tv    int `json:"tv"`
}

type Count struct {
	All      int `json:"all"`
	Rated    int `json:"rated"`
	Ai       int `json:"ai"`
	Draw     int `json:"draw"`
	DrawH    int `json:"drawH"`
	Loss     int `json:"loss"`
	LossH    int `json:"lossH"`
	Win      int `json:"win"`
	WinH     int `json:"winH"`
	Bookmark int `json:"bookmark"`
	Playing  int `json:"playing"`
	Import   int `json:"import"`
	Me       int `json:"me"`
}
