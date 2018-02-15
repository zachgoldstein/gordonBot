package main

import (
	"flag"
	"fmt"

	"github.com/zachgoldstein/gordonBot/slack"
)

func main() {
	apiKey := flag.String("apiKey", "xyz", "API key for slack bot")
	flag.Parse()

	fmt.Println("Starting Gordon. A golang slack bot channelling Gordon Ramsay")

	slackClient := slack.CreateSlackClient(*apiKey)

	slack.RespondToEvents(slackClient)
}
