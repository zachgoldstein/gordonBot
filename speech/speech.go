package speech

import (
	"fmt"
	"strings"
)

var sentenceStartMap = map[string]string{
	"january":   "Way to Go!",
	"february":  "Keep it Up!",
	"march":     "Beautifully Done!",
	"april":     "Killer Moves!",
	"may":       "Cowabungaaaaaa!",
	"june":      "Awesome Job",
	"july":      "Show me what you got!",
	"august":    "I Like Your Moves!",
	"september": "Can't Get Enough!",
	"october":   "Shine On!",
	"november":  "I Believe in You!",
	"december":  "Magnificent Effort!",
}

var sentenceAdjectiveMap = map[string]string{
	"a": "You Rainbow-Infused",
	"b": "You Molten",
	"c": "You Inspiring",
	"d": "You Adorable",
	"e": "You Beautiful",
	"f": "You Poetic",
	"g": "You Talented",
	"h": "You Powerful",
	"i": "You Rule-Breaking",
	"j": "You Brilliant",
	"k": "You Fierce",
	"l": "You Fiery",
	"m": "You Intergalactic",
	"n": "You Highly Intelligent",
	"o": "You Luminescent",
	"p": "You Fluffy",
	"q": "You Kind",
	"r": "You Majestic",
	"s": "You Regal",
	"t": "You Unfathomably Harmonic",
	"u": "You Surprisingly Well-Adjusted",
	"v": "You Softly Glowing",
	"w": "You Melodic",
	"x": "You Dangerously Talented",
	"y": "You Perfectly Spherical",
	"z": "You Clever",
}

var sentenceNounMap = map[string]string{
	"a": "Echidna!",
	"b": "Musk Ox!",
	"c": "Land Mermaid",
	"d": "Narwhal!",
	"e": "Tiger-person!",
	"f": "Kangaroo Rat",
	"g": "Interdimensional Being",
	"h": "Platypus",
	"i": "Combat Wombat",
	"j": "Space Unicorn!",
	"k": "Dolphin!",
	"l": "Marmot!",
	"m": "Bird-person!",
	"n": "Piece of Surprise Glitter!",
	"o": "Manitee!",
	"p": "Remorseless Totoro!",
	"q": "Full-Sized Yoda!",
	"r": "Chimera!",
	"s": "Gigantic Maine Coone Cat",
	"t": "Pomeranian",
	"u": "Champion",
	"v": "Grumpy Cat",
	"w": "Ice Dragon",
	"x": "Toaster",
	"y": "Apocalyptic gladiator",
	"z": "Genius",
}

// CreateMessage will create a message constructed from your birth month, first name and last name.
func CreateMessage(month, firstName, lastName string) (string, error) {
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

// GenerateRandomMessage creates a random message with our sentence mappings.
func GenerateRandomMessage() string {
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

// SplitMessage breaks up a given message into an array expecting three strings
func SplitMessage(message string) ([]string, error) {
	fmt.Println("Message??", message)
	messageParts := strings.Split(message, " ")
	if len(messageParts) != 3 {
		return []string{}, fmt.Errorf("Message is missing information")
	}
	return messageParts, nil
}
