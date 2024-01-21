package dictionary

import (
	"encoding/json"
	"os"
	"slices"
	"strings"
)

const notDefinedMessage string = "Šī frāze netika atrasta vai arī tā vēl nav definēta."

type Phrases struct {
	Phrases []Phrase
}

type Phrase struct {
	Names      []string
	Definition string
}

var phrases Phrases

func Define(name string) string {
	data, err := os.ReadFile("phrases.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &phrases)

	for i := 0; i < len(phrases.Phrases); i++ {
		if slices.Contains(phrases.Phrases[i].Names, strings.ToLower(name)) {
			return phrases.Phrases[i].Definition
		}
	}

	return notDefinedMessage

}
