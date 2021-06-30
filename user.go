package main

type Statuses []UserStatus

type UserStatus struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title,omitempty"`
	Online    bool   `json:"online,omitempty"`
	Playing   bool   `json:"playing,omitempty"`
	Streaming bool   `json:"streaming,omitempty"`
	Patron    bool   `json:"patron,omitempty"`
}

type RatingProgress struct {
	Rating   int `json:"rating"`
	Progress int `json:"progress"`
}

type PublicAccount struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Online   bool   `json:"online,omitempty"`
	Title    string `json:"title"`
	Patron   bool   `json:"patron,omitempty"`
}

type BulletPublicAccount struct {
	Perfs struct {
		Bullet RatingProgress `json:"bullet"`
	} `json:"perfs"`

	PublicAccount
}

type UltraBulletPublicAccount struct {
	Perfs struct {
		UltraBullet RatingProgress `json:"ultrabullet"`
	} `json:"perfs"`

	PublicAccount
}

type BlitzPublicAccount struct {
	Perfs struct {
		Blitz RatingProgress `json:"blitz"`
	} `json:"perfs"`

	PublicAccount
}

type RapidPublicAccount struct {
	Perfs struct {
		Rapid RatingProgress `json:"rapid"`
	} `json:"perfs"`

	PublicAccount
}

type ClassicalPublicAccount struct {
	Perfs struct {
		Classical RatingProgress `json:"classical"`
	} `json:"perfs"`

	PublicAccount
}

type Chess960PublicAccount struct {
	Perfs struct {
		Chess960 RatingProgress `json:"chess960"`
	} `json:"perfs"`

	PublicAccount
}

type CrazyHousePublicAccount struct {
	Perfs struct {
		Crazyhouse RatingProgress `json:"crazyhouse"`
	} `json:"perfs"`

	PublicAccount
}

type AntichessPublicAccount struct {
	Perfs struct {
		Antichess RatingProgress `json:"antichess"`
	} `json:"perfs"`

	PublicAccount
}

type AtomicPublicAccount struct {
	Perfs struct {
		Atomic RatingProgress `json:"atomic"`
	} `json:"perfs"`

	PublicAccount
}

type HordePublicAccount struct {
	Perfs struct {
		Horde RatingProgress `json:"horde"`
	} `json:"perfs"`

	PublicAccount
}

type KingOfTheHillPublicAccount struct {
	Perfs struct {
		KingOfTheHill RatingProgress `json:"kingofthehill"`
	} `json:"perfs"`

	PublicAccount
}

type RacingKingsPublicAccount struct {
	Perfs struct {
		RacingKings RatingProgress `json:"racingkings"`
	} `json:"perfs"`

	PublicAccount
}

type ThreeCheckPublicAccount struct {
	Perfs struct {
		ThreeCheck RatingProgress `json:"threecheck"`
	} `json:"perfs"`

	PublicAccount
}

type TopTenPlayer struct {
	Bullet        []BulletPublicAccount        `json:"bullet"`
	Blitz         []BlitzPublicAccount         `json:"blitz"`
	Rapid         []RapidPublicAccount         `json:"rapid"`
	Classical     []ClassicalPublicAccount     `json:"classical"`
	UltraBullet   []UltraBulletPublicAccount   `json:"ultraBullet"`
	Chess960      []Chess960PublicAccount      `json:"chess960"`
	Crazyhouse    []CrazyHousePublicAccount    `json:"crazyhouse"`
	Antichess     []AntichessPublicAccount     `json:"antichess"`
	Atomic        []AtomicPublicAccount        `json:"atomic"`
	Horde         []HordePublicAccount         `json:"horde"`
	KingOfTheHill []KingOfTheHillPublicAccount `json:"kingOfTheHill"`
	RacingKings   []RacingKingsPublicAccount   `json:"racingKings"`
	ThreeCheck    []ThreeCheckPublicAccount    `json:"threeCheck"`
}

type BlitzLeader struct {
	Users []BlitzPublicAccount `json:"users"`
}

type BulletLeader struct {
	Users []BulletPublicAccount `json:"users"`
}

type UltraBulletLeader struct {
	Users []UltraBulletPublicAccount `json:"users"`
}

type RapidLeader struct {
	Users []RapidPublicAccount `json:"users"`
}

type ClassicalLeader struct {
	Users []ClassicalPublicAccount `json:"users"`
}

type Chess960Leader struct {
	Users []Chess960PublicAccount `json:"users"`
}

type CrazyHouseLeader struct {
	Users []CrazyHousePublicAccount `json:"users"`
}

type AntiChessLeader struct {
	Users []AntichessPublicAccount `json:"users"`
}

type AtomicLeader struct {
	Users []AtomicPublicAccount `json:"users"`
}

type HordeLeader struct {
	Users []HordePublicAccount `json:"users"`
}

type KingOfTheHillLeader struct {
	Users []KingOfTheHillPublicAccount `json:"users"`
}

type RacingKingsLeader struct {
	Users []RacingKingsPublicAccount `json:"users"`
}

type ThreeCheckLeader struct {
	Users []ThreeCheckPublicAccount `json:"users"`
}

type RatingHistory struct {
	Name   string  `json:"name"`
	Points [][]int `json:"points"`
}

type Crosstable struct {
	Users   map[string]float32 `json:"users"`
	NbGames uint32             `json:"nbGames"`
	Matchup struct {
		Users   map[string]float32 `json:"users"`
		NbGames uint32             `json:"nbGames"`
	} `json:"matchup"`
}
