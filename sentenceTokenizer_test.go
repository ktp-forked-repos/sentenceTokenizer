package sentenceTokenizer

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	Init()
	str := "I still haven't heard a good argument as to WHY?  Why should we fix their mortgages.. Keep them in the homes they could not afford?\n\nThere are more responsible people waiting on the sidelines for the prices to drop so they can buy the house.   People who've been waiting because they knew they couldn't afford the absurd prices, and whom are now waiting because prices remain too high.\n\nBailing these people out to keep them in their home only reinforces the mentality that you can do the wrong thing and someone will be there to pat you on the back and say... \"Its ok. Go back into debt, we'll help you again next time\"\n\nIts a big FUCK YOU to the people who were responsible."
	s := Sentences(str)
	for _, sent := range s {
		fmt.Println(sent)
	}
}
