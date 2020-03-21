package main

import (
	"fmt"
	"gopkg.in/neurosnap/sentences.v1"
	"gopkg.in/neurosnap/sentences.v1/data"
	"io/ioutil"
	//"os"
)

func main() {
	//	text := `A perennial also-ran, Stallings won his seat when longtime lawmaker David Holmes
	//    died 11 days after the filing deadline. Suddenly, Stallings was a shoo-in, not
	//  the long shot. In short order, the Legislature attempted to pass a law allowing
	//former U.S. Rep. Carolyn Cheeks Kilpatrick to file; Stallings challenged the
	// law in court and won. Kilpatrick mounted a write-in campaign, but Stallings won.`
	//file, err := os.Open("../our.txt")
	text, _ := ioutil.ReadFile("../our.txt")
	// Compiling language specific data into a binary file can be accomplished
	// by using `make <lang>` and then loading the `json` data:
	b, _ := data.Asset("data/english.json")

	// load the training data
	training, _ := sentences.LoadTraining(b)

	// create the default sentence tokenizer
	tokenizer := sentences.NewSentenceTokenizer(training)
	sentences := tokenizer.Tokenize(string(text))

	for _, s := range sentences {
		fmt.Println(s.Text)
	}
}
