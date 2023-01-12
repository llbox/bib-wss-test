package main

import (
	"bib-wss-test/client"
	"bib-wss-test/message"
)

func main() {
	h := client.NewHub()
	go client.WritePump(h)
	go message.WssSendPump(h)
}
