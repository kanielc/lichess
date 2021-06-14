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
