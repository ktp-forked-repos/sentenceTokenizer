package sentenceTokenizer

import (
	// "fmt"
	"github.com/qubies/stopwords"
	// "gopkg.in/jdkato/prose.v2"
	"gopkg.in/neurosnap/sentences.v1"
	"gopkg.in/neurosnap/sentences.v1/english"
)

var tokenizer *sentences.DefaultSentenceTokenizer

func init() {
	var err error
	tokenizer, err = english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}
}

func Sentences(input string) []string {
	sentences := tokenizer.Tokenize(input)
	list := make([]string, len(sentences))
	for i, sent := range sentences {
		list[i] = stopwords.Remove(sent.Text)
	}
	return list
}
