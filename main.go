package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

var sentenceStartMap = map[string]string{
	"january":   "Fuck Off!",
	"february":  "Shut The Fuck Up!",
	"march":     "Piss Off!",
	"april":     "Get Out!",
	"may":       "It's RAAAW!",
	"june":      "It's Burnt!",
	"july":      "WTF Are You Doing?",
	"august":    "Move Your Ass!",
	"september": "More Sauce!",
	"october":   "Wake Up!",
	"november":  "Get A Grip!",
	"december":  "Gimme Your Jacket!",
}

var sentenceAdjectiveMap = map[string]string{
	"a": "You Stupid",
	"b": "You Lazy",
	"c": "You Pathetic",
	"d": "You Useless",
	"e": "You Silly",
	"f": "You Ingorant",
	"g": "You Fat",
	"h": "You Dumb",
	"i": "You Little",
	"j": "You Fucking",
	"k": "You Bloody",
	"l": "You Ugly",
	"m": "You Weird",
	"n": "You Hopeless",
	"o": "You Wimpy",
	"p": "You Goddamn",
	"q": "You Brainless",
	"r": "You Slow",
	"s": "You Proud",
	"t": "You Fat-mouthed",
	"u": "You Blasted",
	"v": "You Wasted",
	"w": "You Dopey",
	"x": "You Right",
	"y": "You Worthless",
	"z": "You Stinky",
}

var sentenceNounMap = map[string]string{
	"a": "Piece of Shit!",
	"b": "Asshole!",
	"c": "Donut",
	"d": "Idiot!",
	"e": "Jerk!",
	"f": "Pig!",
	"g": "Donkey!",
	"h": "Fuckface!",
	"i": "Wanker!",
	"j": "Cow!",
	"k": "Dumbo!",
	"l": "Imbecile!",
	"m": "Bum!",
	"n": "Muppet",
	"o": "Banana",
	"p": "Dickface",
	"q": "Gremlin",
	"r": "Bozo",
	"s": "Fucker",
	"t": "Fatass!",
	"u": "Dog!",
	"v": "Plank!",
	"w": "Dick!",
	"x": "Giraffe!",
	"y": "Tosser!",
	"z": "CryBaby!",
}

func createInsult(month, firstName, lastName string) (string, error) {
	sentenceParts := []string{}
	sentenceStart, ok := sentenceStartMap[strings.ToLower(month)]
	if !ok {
		return "", fmt.Errorf("Could not find a starting insult with month %s", month)
	}
	sentenceParts = append(sentenceParts, sentenceStart)

	firstNameLetter := strings.ToLower(string(firstName[0]))
	sentenceAdjective, ok := sentenceAdjectiveMap[firstNameLetter]
	if !ok {
		return "", fmt.Errorf("Could not find a starting insult with first name %s", firstName)
	}
	sentenceParts = append(sentenceParts, sentenceAdjective)

	lastNameLetter := strings.ToLower(string(lastName[0]))
	sentenceNoun, ok := sentenceNounMap[lastNameLetter]
	if !ok {
		return "", fmt.Errorf("Could not find a starting insult with last name %s", lastName)
	}
	sentenceParts = append(sentenceParts, sentenceNoun)

	return strings.Join(sentenceParts, " "), nil
}

func generateRandomMessage() string {
	sentenceParts := []string{}
	for _, v := range sentenceStartMap {
		sentenceParts = append(sentenceParts, v)
		break
	}
	for _, v := range sentenceAdjectiveMap {
		sentenceParts = append(sentenceParts, v)
		break
	}
	for _, v := range sentenceNounMap {
		sentenceParts = append(sentenceParts, v)
		break
	}
	return strings.Join(sentenceParts, " ")
}

func splitMessage(message string) ([]string, error) {
	messageParts := strings.Split(message, " ")
	if len(messageParts) < 3 {
		return []string{}, fmt.Errorf("Message is missing information")
	}
	return messageParts, nil
}

func main() {
	api := slack.New("XYZ")
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)
			fmt.Printf("Message text %s\n", ev.Msg.Text)
			fmt.Printf("CHANNEL ID: %v\n", ev.Channel)
			if !strings.Contains(ev.Msg.Text, "<@U94LXJJPM>") {
				continue
			}
			message := strings.Replace(ev.Msg.Text, "<@U94LXJJPM> ", "", -1)
			if strings.ToLower(message) == "help" {
				rtm.SendMessage(rtm.NewOutgoingMessage("type in '@gordon <month of birth> <firstname> <lastname>'", ev.Channel))
				continue
			}

			sentenceParts, err := splitMessage(message)
			if err != nil {
				fmt.Printf("Error splitting message: %s\n", err)
				rtm.SendMessage(rtm.NewOutgoingMessage(generateRandomMessage(), ev.Channel))
				continue
			}
			insult, err := createInsult(sentenceParts[0], sentenceParts[1], sentenceParts[2])
			if err != nil {
				fmt.Printf("Error creating insult: %s\n", err)
				rtm.SendMessage(rtm.NewOutgoingMessage(generateRandomMessage(), ev.Channel))
				continue
			}
			rtm.SendMessage(rtm.NewOutgoingMessage(insult, ev.Channel))

		case *slack.PresenceChangeEvent:
			fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:

			// Ignore other events..
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
