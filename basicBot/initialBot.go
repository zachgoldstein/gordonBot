package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("XXYY")
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
	}

}
