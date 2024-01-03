package main

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"time"
	"io/ioutil"
	"strconv"
)

type Credentials struct {
	Token, Author string
	Channels []string
}

type Message struct {
	Id string
	Type int
	Channel string `json:"channel_id"`
}

type Messages struct {
	Total int `json:"total_results"`
	Messages [][]Message
}

func main() {
	stack := map[string][]string{}

	p, err := os.ReadFile("preset.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var preset Credentials 
	err = json.Unmarshal(p, &preset)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}

	for _, channel := range preset.Channels {
		target := "channels"	
		offset := 0
		totalMessages := 1
		for offset < totalMessages {
			req, err := http.NewRequest("GET", 
			"https://discord.com/api/v9/"+target+"/"+channel+
			"/messages/search?author_id="+preset.Author+
			"&offset="+strconv.Itoa(offset), 
			nil)

			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Add("Authorization", preset.Token)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}

			if resp.StatusCode == 404 {
				fmt.Println("not found in", target)
				if target == "channels" {
					target = "guilds"
				} else {
					target = "channels"
				}
				fmt.Println("look for", target)
				continue
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			var messages Messages
			err = json.Unmarshal(body, &messages)

			if resp.StatusCode == 200 {
				offset += len(messages.Messages) 
				totalMessages = messages.Total
			} else {
				fmt.Println("not expected response", resp.StatusCode, "try again")
				continue
			}

			fmt.Println("fetching messages", offset, "/", totalMessages, "from channel", channel, resp.StatusCode)

			for _, message := range messages.Messages {
				if message[0].Type == 0 || message[0].Type == 19 {
					stack[message[0].Channel] = append(
						stack[message[0].Channel], 
						message[0].Id)
				}
			}

			resp.Body.Close()
			time.Sleep(5 * time.Second)
		}
	}

	i := 0
	for channel, messages := range stack {
		t := 0
		for t < len(messages) {
			req, err := http.NewRequest("DELETE", 
			"https://discord.com/api/v9/channels/"+channel+
			"/messages/"+messages[t], nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			req.Header.Add("Authorization", preset.Token)
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("deleting message ", (t+1), "/", len(messages), messages[t], "from channel", (i+1), "/", len(stack), channel, resp.StatusCode)
			if resp.StatusCode == 204 {
				t += 1
			}
			resp.Body.Close()
			time.Sleep(3 * time.Second)
		}

		i += 1
	}
	return
}
