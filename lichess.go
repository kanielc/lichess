package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

// useful for testing the calls
func main() {
	token := flag.String("token", "", "Access Token")
	flag.Parse()

	if *token == "" {
		flag.Usage()
		os.Exit(0)
	}

	httpClient := http.Client{
		Timeout: 3 * time.Second,
	}

	client := Client{
		Token:      *token,
		HttpClient: &httpClient,
	}

	acct, _ := client.FetchAccount()

	//bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyBytes))
	fmt.Println(acct.Language)
	fmt.Println(acct.Perfs.Blitz.Rating)
	fmt.Println(acct.Email)
}
