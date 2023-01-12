package main

import (
	"bib-wss-test/client"
	"bib-wss-test/service"
)

func main() {
	h := client.NewHub()
	h.Run()
	go client.WritePump(h)
	go service.OrderPush(h)
}
