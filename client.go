package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	LichessBase = "https://lichess.org"
)

type Client struct {
	Token      string
	HttpClient *http.Client
}

func (c *Client) NewRequest(method string, url string, body *io.Reader) (*http.Request, error) {

	var reader io.Reader
	if body == nil {
		reader = new(bytes.Buffer)
	} else {
		reader = *body
	}

	req, err := http.NewRequest(method, url, reader)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Accept", "application/json")
	return req, err
}

func (c *Client) doGet(endPoint string, dest interface{}) (*http.Response, error) {
	req, err := c.NewRequest("GET", LichessBase+"/api/account", nil)
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

func (c *Client) FetchEmail() (string, error) {
	var email Email
	resp, err := c.doGet("/api/account/email", &email)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return email.Email, nil
}

func (c *Client) FetchAccount() (*Account, error) {
	var acct Account
	resp, err := c.doGet("/api/account", &acct)
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
