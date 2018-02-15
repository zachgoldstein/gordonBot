package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

		switch ev := msg.Data.(type) {

		case *slack.MessageEvent:
			botTagString := fmt.Sprintf("<@%s>", rtm.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			rtm.SendMessage(rtm.NewOutgoingMessage("IT'S BURNTTTT!", ev.Channel))
		default:

		}
	}
}
