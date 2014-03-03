package bay

import (
	"fmt"
	"testing"
)

func Test_NBC(test *testing.T) {
	// 1. Read/Import training data (DATA) , from my GitHub / Google Docs
	DATA_amazon := read.GetStruct("train - amazon.csv")
	DATA_city := read.GetStruct("train - city.csv")
	include := read.GetInclFt("filter - include.csv")
	exclude := read.GetExcFt("filter - exclude.csv")

	// Pass unfamiliar sentences and see how accurate its sentiment analysis is.
	str1 := "I highly recommend here. Great Weather!"
	str2 := "I hate the movie."
	str3 := "I enjoy it and want to do this again."
	str4 := "Quite disappointed. Never ever again!"
	str5 := "And all of this at a great price."

	fmt.Println("Bayesian Sentiment Analysis in Amazon.com model.")
	Print(DATA_amazon, include, exclude, str1)
	Print(DATA_amazon, include, exclude, str2)
	Print(DATA_amazon, include, exclude, str3)
	Print(DATA_amazon, include, exclude, str4)
	Print(DATA_amazon, include, exclude, str5)

	/*
		Bayesian Sentiment Analysis in Amazon.com model.
		Strongly Positive: I highly recommend here. Great Weather!
		Strongly Negative: I hate this movie.
		Strongly Positive: I enjoy it and want to do this again.
		Negative: Quite disappointed. Never ever again!
		Strongly Positive: And all of this at a great price.
	*/

	println()
	println()

	fmt.Println("Bayesian Sentiment Analysis in city review model.")
	Print(DATA_city, include, exclude, str1)
	Print(DATA_city, include, exclude, str2)
	Print(DATA_city, include, exclude, str3)
	Print(DATA_city, include, exclude, str4)
	Print(DATA_city, include, exclude, str5)

	/*
		Bayesian Sentiment Analysis in city review model.
		Positive: I highly recommend here. Great Weather!
		Strongly Negative: I hate this movie.
		Strongly Positive: I enjoy it and want to do this again.
		Strongly Negative: Quite disappointed. Never ever again!
		Strongly Positive: And all of this at a great price.
	*/

	// More Example
	Print(DATA_amazon, include, exclude, "High quality code samples. It must be said: Mark Summerfield is a REALLY good programmer. All of the code in this book gives the impression of being well thought out. The other books had a lot of cargo cult programming, meaning the authors were going through the motions without thinking about what they were doing.")
	// Totally unfamiliar sentence
	// Positive: ~ (Correct Classification!)
	// Now this data is trained

	Print(DATA_amazon, include, exclude, "I just paid good money for this book and went to the web site to download the code examples and exercises. The web page is almost totally unreadable since it has a black background with dark gray text. But the worst thing is that there are no links to download the code examples and exercises. While I can read the 3 chapters of updated text in my browser (chrome), the save button doesn't work, I can only print them. There is no link to tell the author any of this, so I have to do it here. I am new to Kindle, but I don't see how I can put this new material back into my Kindle book.")
	// Totally unfamiliar sentence
	// Negative: ~ (Correct Classification!)
	// Now this data is trained

	Print(DATA_amazon, include, exclude, "The book is full of great content and contains the most comprehensive treatment on Go channels and goroutines that I've seen yet. The author also does a phenomenal job with coverage; the table of contents shows this. Virtually every nook and cranny of the Go language is covered. Including the initial founding of Go and its context. (I found this last bit very interesting.) The book is also full of code snippets exemplifying Go idioms and as examples to accomplish common tasks, which is really great considering the infancy of Go.")
	// Totally unfamiliar sentence
	// Positive: ~ (Correct Classification!)
	// Now this data is trained
}
