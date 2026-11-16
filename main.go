package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2/clientcredentials"
)

type searchResp struct {
	Artists struct {
		Items []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			URL  struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
		} `json:"items"`
	} `json:"artists"`
}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	ctx := context.Background()
	cfg := clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     "https://accounts.spotify.com/api/token",
	}

	client := cfg.Client(ctx)

	q := url.Values{
		"q": {"Daft Punk"}, "type": {"artist"}, "limit": {"3"},
	}

	resp, err := client.Get("https://api.spotify.com/v1/search?" + q.Encode())

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	var out searchResp
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		log.Fatal(err)
	}
	for _, a := range out.Artists.Items {
		fmt.Printf("%-40s %s\n", a.Name, a.URL.Spotify)
	}
}
