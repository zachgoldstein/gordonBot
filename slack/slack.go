package slack

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
	"github.com/zachgoldstein/gordonBot/speech"
)

const helpMessage = "type in '@gordon <month of birth> <firstname> <lastname>'"

// CreateSlackClient sets up the slack RTM (real-timemessaging) client library,
// initiating the socket connection and returning the client
func CreateSlackClient(apiKey string) *slack.RTM {
	api := slack.New(apiKey)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()
	return rtm
}

// RespondToEvents waits for messages on the slack client's incomingEvents
// channel, sending responses when it detects the bot has been tagged in a
// message with @<botTag>
func RespondToEvents(slackClient *slack.RTM) {
	for msg := range slackClient.IncomingEvents {
		fmt.Println("Event Received: ", msg.Type)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			botTagString := fmt.Sprintf("<@%s> ", slackClient.GetInfo().User.ID)
			if !strings.Contains(ev.Msg.Text, botTagString) {
				continue
			}
			message := strings.Replace(ev.Msg.Text, botTagString, "", -1)

			sendHelp(slackClient, message, ev.Channel)

			sendResponse(slackClient, message, ev.Channel)

		default:

		}
	}
}

func sendHelp(slackClient *slack.RTM, message, slackChannel string) {
	if strings.ToLower(message) != "help" {
		return
	}
	slackClient.SendMessage(slackClient.NewOutgoingMessage(helpMessage, slackChannel))
}

func sendResponse(slackClient *slack.RTM, message, slackChannel string) {
	sentenceParts, err := speech.SplitMessage(message)
	if err != nil {
		fmt.Printf("Error splitting message: %s\n", err)
		slackClient.SendMessage(slackClient.NewOutgoingMessage(speech.GenerateRandomMessage(), slackChannel))
		return
	}

	response, err := speech.CreateMessage(sentenceParts[0], sentenceParts[1], sentenceParts[2])
	if err != nil {
		fmt.Printf("Error creating response: %s\n", err)
		slackClient.SendMessage(slackClient.NewOutgoingMessage(speech.GenerateRandomMessage(), slackChannel))
		return
	}
	slackClient.SendMessage(slackClient.NewOutgoingMessage(response, slackChannel))
}
