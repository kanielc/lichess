package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	LichessBase = "https://lichess.org"
)

type Client struct {
	Token      string
	HttpClient *http.Client
}

type RequestParams struct {
	Method        string
	Accept        string
	ContentType   string
	Body          *io.Reader
	Authorization string
	QueryValues   url.Values
}

func (c *Client) DefaultRequestParams() *RequestParams {
	return &RequestParams{
		Method:        "GET",
		Accept:        "application/json",
		ContentType:   "",
		Authorization: fmt.Sprintf("Bearer %s", c.Token),
		QueryValues:   url.Values{},
	}
}

func (c *Client) NewRequest(url string, params *RequestParams) (*http.Request, error) {
	if params == nil {
		params = c.DefaultRequestParams()
	}

	var conts io.Reader
	if params.Body != nil {
		conts = *params.Body
	} else {
		conts = strings.NewReader(params.QueryValues.Encode())
	}

	req, err := http.NewRequest(params.Method, url, conts)
	if err != nil {
		return nil, err
	}

	if params.Authorization != "" {
		req.Header.Set("Authorization", params.Authorization)
	}

	if params.Accept != "" {
		req.Header.Set("Accept", params.Accept)
	}

	if params.ContentType != "" {
		req.Header.Set("Content-Type", params.ContentType)
	}

	return req, err
}

func (c *Client) DoRequest(endPoint string, dest interface{}, params *RequestParams) (*http.Response, error) {
	req, err := c.NewRequest(LichessBase+endPoint, params)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(dest)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) GetEmail() (string, error) {
	var email Email
	resp, err := c.DoRequest("/api/account/email", &email, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return email.Email, nil
}

func (c *Client) GetAccount() (*Account, error) {
	var acct Account
	resp, err := c.DoRequest("/api/account", &acct, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	email, err := c.GetEmail()

	if err != nil || acct.Email == "" {
		acct.Email = "N/A"
	} else {
		acct.Email = email
	}

	return &acct, nil
}

func (c *Client) GetUser(id string) (*Account, error) {
	var acct Account

	resp, err := c.DoRequest("/api/user/"+id, &acct, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &acct, nil
}

func (c *Client) GetUsers(ids ...string) ([]Account, error) {
	var accts = make([]Account, 0)
	allIds := strings.Join(ids, ",")
	params := c.DefaultRequestParams()
	params.Method = "POST"
	var read io.Reader = bytes.NewReader([]byte(allIds))
	params.Body = &read

	resp, err := c.DoRequest("/api/users", &accts, params)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return accts, nil
}

func (c *Client) GetUserStatus(users []string) ([]UserStatus, error) {
	if len(users) == 0 {
		return nil, errors.New("no users provided, cannot be nil or empty")
	}
	statuses := make([]UserStatus, 0)

	var ids string
	if len(users) > 0 {
		ids = "?ids=" + strings.Join(users, ",")
	}
	resp, err := c.DoRequest("/api/users/status"+ids, &statuses, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return statuses, nil
}

func (c *Client) GetTeamMembers(teamId string) ([]Account, error) {
	if teamId == "" {
		return nil, errors.New("no valid teamId provided, cannot be empty")
	}

	params := c.DefaultRequestParams()
	params.Accept = "application/x-ndjson"
	params.Authorization = ""
	team := make([]Account, 0)

	uri := LichessBase + "/api/team/" + teamId + "/users"
	req, err := c.NewRequest(uri, params)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var acct Account
	decoder := json.NewDecoder(resp.Body)
	for decoder.More() {
		err = decoder.Decode(&acct)

		if err != nil {
			return nil, err
		}
		team = append(team, acct)
	}
	defer resp.Body.Close()

	return team, nil
}

func (c *Client) GetTopTenPlayers() (*TopTenPlayer, error) {
	var topTen TopTenPlayer
	params := c.DefaultRequestParams()
	params.Accept = "application/vnd.lichess.v3+json"
	resp, err := c.DoRequest("/player", &topTen, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &topTen, nil
}

func (c *Client) GetLeaderBoard(number int, gameType string) (interface{}, error) {
	var leaderType interface{}
	var resp *http.Response
	var err error

	params := c.DefaultRequestParams()
	params.Accept = "application/vnd.lichess.v3+json"

	endPoint := fmt.Sprintf("/player/top/%d/%s", number, gameType)
	switch gameType {
	case "blitz":
		var blitzLeader BlitzLeader
		resp, err = c.DoRequest(endPoint, &blitzLeader, params)
		leaderType = blitzLeader
	case "bullet":
		var bulletLeader BulletLeader
		resp, err = c.DoRequest(endPoint, &bulletLeader, params)
		leaderType = bulletLeader
	case "ultraBullet":
		var leader UltraBulletLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "rapid":
		var leader RapidLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "classical":
		var leader ClassicalLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "chess960":
		var leader Chess960Leader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "crazyhouse":
		var leader CrazyHouseLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "antichess":
		var leader AntiChessLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "atomic":
		var leader AtomicLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "horde":
		var leader HordeLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "kingOfTheHill":
		var leader KingOfTheHillLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "racingKings":
		var leader RacingKingsLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	case "threeCheck":
		var leader ThreeCheckLeader
		resp, err = c.DoRequest(endPoint, &leader, params)
		leaderType = leader
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return leaderType, nil
}

func (c *Client) GetRatingHistory(id string) ([]RatingHistory, error) {
	var ratingHistory = make([]RatingHistory, 0)

	resp, err := c.DoRequest("/api/user/"+id+"/rating-history", &ratingHistory, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ratingHistory, nil
}

func (c *Client) GetLiveStreamers() ([]BasicAccount, error) {
	var accts = make([]BasicAccount, 0)

	resp, err := c.DoRequest("/streamer/live", &accts, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return accts, nil
}

func (c *Client) GetCrosstable(user1, user2 string) (*Crosstable, error) {
	if user1 == "" || user2 == "" {
		return nil, fmt.Errorf("user names must be valid and of non-nil length, given user1: %s and user2: %s", user1, user2)
	}
	var crosstable Crosstable

	resp, err := c.DoRequest(fmt.Sprintf("/api/crosstable/%s/%s", user1, user2), &crosstable, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &crosstable, nil
}

func (c *Client) GetFollows(user string) ([]Account, error) {
	if user == "" {
		return nil, errors.New("must have valid id to check follows of")
	}

	uri := LichessBase + "/api/user/" + user + "/following"
	params := c.DefaultRequestParams()
	params.Accept = "application/x-ndjson"

	req, err := c.NewRequest(uri, params)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var follows = make([]Account, 0)
	var acct Account
	decoder := json.NewDecoder(resp.Body)
	for decoder.More() {
		err = decoder.Decode(&acct)

		if err != nil {
			return nil, err
		}
		follows = append(follows, acct)
	}
	defer resp.Body.Close()

	return follows, nil
}

func (c *Client) GetFollowers(user string) ([]Account, error) {
	if user == "" {
		return nil, errors.New("must have valid id to check follows of")
	}

	uri := LichessBase + "/api/user/" + user + "/followers"
	params := c.DefaultRequestParams()
	params.Accept = "application/x-ndjson"

	req, err := c.NewRequest(uri, params)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var followers = make([]Account, 0)
	var acct Account
	decoder := json.NewDecoder(resp.Body)
	for decoder.More() {
		err = decoder.Decode(&acct)

		if err != nil {
			return nil, err
		}
		followers = append(followers, acct)
	}
	defer resp.Body.Close()

	return followers, nil
}

func (c *Client) GetGame(gameId string, params GameParam) (*Game, error) {
	if gameId == "" {
		return nil, errors.New("must provide a valid game id")
	}

	var game Game

	resp, err := c.DoRequest(fmt.Sprintf("/game/export/%s", gameId), &game, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &game, nil
}
