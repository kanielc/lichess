package main

import (
	"encoding/json"
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
	Body          *io.Reader
	Authorization string
	QueryValues   url.Values
}

func (c *Client) DefaultRequestParams() *RequestParams {
	return &RequestParams{
		Method:        "GET",
		Accept:        "application/json",
		Authorization: fmt.Sprintf("Bearer %s", c.Token),
		QueryValues:   url.Values{},
	}
}

func (c *Client) NewRequest(url string, params *RequestParams) (*http.Request, error) {
	if params == nil {
		params = c.DefaultRequestParams()
	}

	req, err := http.NewRequest(params.Method, url, strings.NewReader(params.QueryValues.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", params.Authorization)
	req.Header.Set("Accept", params.Accept)

	return req, err
}

func (c *Client) doGet(endPoint string, dest interface{}, params *RequestParams) (*http.Response, error) {
	req, err := c.NewRequest(LichessBase+endPoint, params)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	//fmt.Println(req.Header)
	//bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyBytes))
	err = json.NewDecoder(resp.Body).Decode(dest)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) FetchEmail() (string, error) {
	var email Email
	resp, err := c.doGet("/api/account/email", &email, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return email.Email, nil
}

func (c *Client) FetchAccount() (*Account, error) {
	var acct Account
	resp, err := c.doGet("/api/account", &acct, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	email, err := c.FetchEmail()

	if err != nil || acct.Email == "" {
		acct.Email = "N/A"
	} else {
		acct.Email = email
	}

	return &acct, nil
}

func (c *Client) FetchUserStatus(users []string) ([]UserStatus, error) {
	statuses := make([]UserStatus, 0)

	var ids string
	if len(users) > 0 {
		ids = "?ids=" + strings.Join(users, ",")
	}
	resp, err := c.doGet("/api/users/status"+ids, &statuses, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return statuses, nil
}

func (c *Client) GetTopTenPlayers() (*TopTenPlayer, error) {
	var topTen TopTenPlayer
	params := c.DefaultRequestParams()
	params.Accept = "application/vnd.lichess.v3+json"
	resp, err := c.doGet("/player", &topTen, params)
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
		resp, err = c.doGet(endPoint, &blitzLeader, params)
		leaderType = blitzLeader
	case "bullet":
		var bulletLeader BulletLeader
		resp, err = c.doGet(endPoint, &bulletLeader, params)
		leaderType = bulletLeader
	case "ultraBullet":
		var leader UltraBulletLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "rapid":
		var leader RapidLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "classical":
		var leader ClassicalLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "chess960":
		var leader Chess960Leader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "crazyhouse":
		var leader CrazyHouseLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "antichess":
		var leader AntiChessLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "atomic":
		var leader AtomicLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "horde":
		var leader HordeLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "kingOfTheHill":
		var leader KingOfTheHillLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "racingKings":
		var leader RacingKingsLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	case "threeCheck":
		var leader ThreeCheckLeader
		resp, err = c.doGet(endPoint, &leader, params)
		leaderType = leader
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	return leaderType, nil
}
