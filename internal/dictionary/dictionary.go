package dictionary

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"slices"
	"strings"
)

const notDefinedMessage string = "Šī frāze netika atrasta vai arī tā vēl nav definēta."

var notDefined = Phrase{
	Names:      []string{notDefinedMessage},
	Definition: "",
	Usage:      "",
}

type Phrases struct {
	Phrases []Phrase
}

type Phrase struct {
	Names      []string
	Definition string
	Usage      string
}

var phrases Phrases

func Define(name string) Phrase {
	data, err := os.ReadFile("phrases.json")
	if err != nil {
		log.Fatal("Error while reading phrases file: ", err)
	}

	json.Unmarshal(data, &phrases)

	for i := 0; i < len(phrases.Phrases); i++ {
		if slices.Contains(phrases.Phrases[i].Names, strings.ToLower(name)) {
			return phrases.Phrases[i]
		}
	}

	return notDefined
}

func Stats() int {
	data, err := os.ReadFile("phrases.json")
	if err != nil {
		log.Fatal("Error while reading phrases file: ", err)
	}

	json.Unmarshal(data, &phrases)
	return len(phrases.Phrases)
}

func Random() Phrase {
	data, err := os.ReadFile("phrases.json")
	if err != nil {
		log.Fatal("Error while reading phrases file: ", err)
	}

	json.Unmarshal(data, &phrases)

	return phrases.Phrases[rand.Intn(len(phrases.Phrases))]
}
