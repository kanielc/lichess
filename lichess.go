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

	//acct, _ := client.FetchAccount()

	//bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bodyBytes))
	/*
			fmt.Println(acct.Language)
			fmt.Println(acct.Perfs.Blitz.Rating)
			fmt.Println(acct.Email)

			if users, err := client.FetchUserStatus([]string{"chess-network", "STL_Nakamura"}); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(users)
			}

			if topTen, err := client.GetTopTenPlayers(); err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("%+v", topTen)
			}

		if leader, err := client.GetLeaderBoard(10, "horde"); err != nil {
			fmt.Println(err)
		} else {
			str := leader.(HordeLeader)
			fmt.Println(str.Users[0].ID, str.Users[0].Username)
		}*/

	if player, err := client.GetUser("chess-network"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(player.ID, player.Online)
	}

	if playerHistory, err := client.GetRatingHistory("chess-network"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(playerHistory[0].Name, playerHistory[0].Points[0])
	}
}
