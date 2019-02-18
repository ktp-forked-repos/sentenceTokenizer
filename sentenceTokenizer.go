package sentenceTokenizer

import (
	"fmt"
	"github.com/qubies/stopwords"
	"gopkg.in/jdkato/prose.v2"
)

func Sentences(input string) []string {
	doc, err := prose.NewDocument(
		input,
		prose.WithExtraction(false), prose.WithTagging(false), prose.WithTokenization(false))
	if err != nil {
		fmt.Println(err)
	}
	var list []string
	for _, sent := range doc.Sentences() {
		list = append(list, stopwords.Remove(sent.Text))
	}
	return list
}
