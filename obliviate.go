package main

import (
	"fmt"
	"os"
	"encoding/json"
	"strconv"
	"net/http"
	"io/ioutil"
	"time"
)

func main() {
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
		offset := 0
		totalMessages := 1
		for offset < totalMessages {
			req, err := http.NewRequest("GET", 
			"https://discord.com/api/v9/channels/"+channel+
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
			}

			fmt.Println("fetching messages "+
			strconv.Itoa(offset)+"/"+strconv.Itoa(totalMessages)+
			" from channel "+channel, resp.StatusCode)

			for _, message := range messages.Messages {
				if message[0].Type == 0 {
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
			fmt.Println("deleting message "+strconv.Itoa(t+1)+"/"+strconv.Itoa(len(messages))+" "+messages[t]+
			" from channel "+strconv.Itoa(i+1)+"/"+strconv.Itoa(len(stack))+" "+channel, resp.StatusCode)
			if resp.StatusCode == 204 {
				t += 1
			}
			resp.Body.Close()
			time.Sleep(3 * time.Second)
		}

		i += i
	}
	return
}
