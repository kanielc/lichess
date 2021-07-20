package main

import (
	"net/http"
	"os"
	"testing"
	"time"
)

var httpClient = http.Client{
	Timeout: 3 * time.Second,
}

var client = Client{
	Token:      os.Getenv("LICHESS_TOKEN"),
	HttpClient: &httpClient,
}

func TestGetAccount(t *testing.T) {
	acct, _ := client.GetAccount()

	if acct.Language == "" {
		t.Errorf("Language is missing")
	}

	if acct.Perfs.Blitz.Rating == 0 && acct.Perfs.Rapid.Rating == 0 && acct.Perfs.Bullet.Rating == 0 {
		t.Errorf("Expected a rating in either Blitz, Rapid or Bullet")
	}

	if acct.Email == "" {
		t.Errorf("Email should be present or N/A")
	}
}

func TestGetUserStatus(t *testing.T) {
	users, _ := client.GetUserStatus([]string{"chess-network", "STL_Nakamura"})

	if users[0].ID == "" || users[1].Name == "" {
		t.Errorf("Expected to get user ID and Name users0: %+v  users[1]: %+v", users[0], users[1])
	}
}

func TestGetTopTenPlayers(t *testing.T) {
	topTen, _ := client.GetTopTenPlayers()

	if topTen.Blitz[3].Perfs.Blitz.Rating == 0 {
		t.Errorf("Expected to get a blitz rating for this top ten player")
	}
}

func TestGetLeaderBoard(t *testing.T) {
	leader, _ := client.GetLeaderBoard(10, "horde")
	str := leader.(HordeLeader)
	user1 := str.Users[0]

	if user1.ID == "" {
		t.Errorf("Invalid User ID from Leaderboard")
	}

	if user1.Username == "" {
		t.Errorf("Invalid Username from Leaderboard")
	}
}

func TestGetUser(t *testing.T) {
	player, _ := client.GetUser("chess-network")

	if player.ID == "" {
		t.Errorf("Expected to get User ID")
	}
}

func TestGetUsers(t *testing.T) {
	players, _ := client.GetUsers("chess-network", "STL_Nakamura")

	if players[0].ID == "" || players[1].ID == "" {
		t.Errorf("Expected to get User ID, instead got user1: +%v and user2: +%v", players[0], players[1])
	}
}

func TestGetRatingHistory(t *testing.T) {
	playerHistory, _ := client.GetRatingHistory("chess-network")

	if playerHistory[0].Name == "" {
		t.Errorf("Expected to get player history (and name)")
	}
}

func TestGetTeamMembers(t *testing.T) {
	team, _ := client.GetTeamMembers("coders")

	if team[0].ID == "" || team[1].ID == "" || team[2].ID == "" {
		t.Errorf("Expected valid team member IDs")
	}
}

func TestGetLiveStreamers(t *testing.T) {
	streamers, _ := client.GetLiveStreamers()

	if streamers[0].ID == "" || streamers[1].ID == "" || streamers[2].ID == "" {
		t.Errorf("Expected valid streamer IDs")
	}
}

func TestGetCrosstable(t *testing.T) {
	crosstable, _ := client.GetCrosstable("neio", "thibault")

	if _, ok := crosstable.Users["neio"]; !ok {
		t.Errorf("Expected to find given user in the crosstable")
	}
}

func TestGetFollows(t *testing.T) {
	follows, _ := client.GetFollows("thibault")

	if follows[0].ID == "" {
		t.Errorf("Expected to find follow")
	}

	if len(follows) < 10 {
		t.Errorf("Expected to find at least 10 follows for thibault")
	}
}

func TestGetFollowers(t *testing.T) {
	followers, _ := client.GetFollowers("thibault")

	if followers[0].ID == "" {
		t.Errorf("Expected to find follow")
	}

	if len(followers) < 2 {
		t.Errorf("Expected to find at least 2 followers for thibault")
	}
}

func TestGetGame(t *testing.T) {
	game, _ := client.GetGame("XWWk5HG6", NewGameParam())

	if game.ID == "" {
		t.Errorf("Expected a game id in the response")
	}

	if game.Moves == "" {
		t.Errorf("Didn't get the moves from the game")
	}
}
