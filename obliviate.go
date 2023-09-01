package main

import ( 
	"net/http"
	"io"
	"fmt"
	"encoding/json"
	"time"
	"strconv"
	"os"
)

type Message struct {
	Id string
	Content string 
	Channel string `json:"channel_id"`
}

type Search struct {
	TotalResults int `json:"total_results"`
	Messages [][]Message
}

type Preset struct {
	Token string
	Author string
	Channel string
}

var p Preset

func actionRequest(method string, url string) ([]byte, int) {
	req, _ := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", p.Token)
	client := &http.Client{}
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	return body, resp.StatusCode
}

func getMessages(offset int) Search {
	body, statusCode := actionRequest("GET", "https://discord.com/api/v9/channels/"+p.Channel+"/messages/search?author_id="+p.Author+"&offset="+strconv.Itoa(offset))

	if statusCode == 404 {
		body, statusCode = actionRequest("GET", "https://discord.com/api/v9/guilds/"+p.Channel+"/messages/search?author_id="+p.Author+"&offset="+strconv.Itoa(offset))
	}

	var s Search
	json.Unmarshal(body, &s)

	return s
}

func main() {
	preset, _ := os.ReadFile("preset.json")
	json.Unmarshal(preset, &p)

	var offset int = 0
	var totalResults int = 9999999

	for totalResults != offset {
		s := getMessages(offset)
		totalResults = s.TotalResults

		if offset == totalResults {
			break
		}

		for _, msg := range s.Messages {
			_, statusCode := actionRequest("DELETE", "https://discord.com/api/v9/channels/"+msg[0].Channel+"/messages/"+msg[0].Id)

			if statusCode == 403 {
				offset++
			}

			time.Sleep(3 * time.Second)

			fmt.Println(msg[0].Id, statusCode)
		}
	}
}
