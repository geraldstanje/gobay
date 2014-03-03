package example

import (
	"fmt"
	"testing"

	"github.com/gyuho/gobay/bay"
	// go test -v github.com/gyuho/gobay/example
	// go test -v /Users/gyuho/go/src/github.com/gyuho/gobay/example/bay_test.go
)

func Test_NBC_0(test *testing.T) {
	// 1. Read/Import training data (DATA) , from my GitHub / Google Docs
	DATA_amazon := bay.GetStruct("../data/train - amazon.csv")
	include := bay.GetInclFt("../data/filter - include.csv")
	exclude := bay.GetExcFt("../data/filter - exclude.csv")

	// Pass unfamiliar sentences and see how accurate its sentiment analysis is.
	str1 := "I highly recommend here. Great Weather!"
	str2 := "I hate the movie."
	str3 := "I enjoy it and want to do this again."
	str4 := "Quite disappointed. Never ever again!"
	str5 := "And all of this at a great price."

	fmt.Println("Bayesian Sentiment Analysis in Amazon.com model.")
	bay.Print(DATA_amazon, include, exclude, str1)
	bay.Print(DATA_amazon, include, exclude, str2)
	bay.Print(DATA_amazon, include, exclude, str3)
	bay.Print(DATA_amazon, include, exclude, str4)
	bay.Print(DATA_amazon, include, exclude, str5)

	// Output:
	// Bayesian Sentiment Analysis in Amazon.com model.
	// Strongly Positive: I highly recommend here. Great Weather!
	// Strongly Negative: I hate this movie.
	// Strongly Positive: I enjoy it and want to do this again.
	// Negative: Quite disappointed. Never ever again!
	// Strongly Positive: And all of this at a great price.
}

func Test_NBC_1(test *testing.T) {
	DATA_city := bay.GetStruct("../data/train - city.csv")
	include := bay.GetInclFt("../data/filter - include.csv")
	exclude := bay.GetExcFt("../data/filter - exclude.csv")

	// Pass unfamiliar sentences and see how accurate its sentiment analysis is.
	str1 := "I highly recommend here. Great Weather!"
	str2 := "I hate the movie."
	str3 := "I enjoy it and want to do this again."
	str4 := "Quite disappointed. Never ever again!"
	str5 := "And all of this at a great price."

	fmt.Println("Bayesian Sentiment Analysis in city review model.")
	bay.Print(DATA_city, include, exclude, str1)
	bay.Print(DATA_city, include, exclude, str2)
	bay.Print(DATA_city, include, exclude, str3)
	bay.Print(DATA_city, include, exclude, str4)
	bay.Print(DATA_city, include, exclude, str5)

	// Output:
	// Bayesian Sentiment Analysis in city review model.
	// Positive: I highly recommend here. Great Weather!
	// Strongly Negative: I hate this movie.
	// Strongly Positive: I enjoy it and want to do this again.
	// Strongly Negative: Quite disappointed. Never ever again!
	// Strongly Positive: And all of this at a great price.
}

func Test_NBC_2(test *testing.T) {
	DATA_amazon := bay.GetStruct("../data/train - amazon.csv")
	include := bay.GetInclFt("../data/filter - include.csv")
	exclude := bay.GetExcFt("../data/filter - exclude.csv")

	// Totally unfamiliar sentence
	// (Correct Classification!)
	// Now this data is trained

	bay.Print(DATA_amazon, include, exclude, "High quality code samples. It must be said: Mark Summerfield is a REALLY good programmer. All of the code in this book gives the impression of being well thought out. The other books had a lot of cargo cult programming, meaning the authors were going through the motions without thinking about what they were doing.")

	// Output:
	// Positive: High quality code samples. It must be said: Mark Summerfield is a REALLY good programmer. All of the code in this book gives the impression of being well thought out. The other books had a lot of cargo cult programming, meaning the authors were going through the motions without thinking about what they were doing.
}

func Test_NBC_3(test *testing.T) {
	DATA_amazon := bay.GetStruct("../data/train - amazon.csv")
	include := bay.GetInclFt("../data/filter - include.csv")
	exclude := bay.GetExcFt("../data/filter - exclude.csv")

	bay.Print(DATA_amazon, include, exclude, "I just paid good money for this book and went to the web site to download the code examples and exercises. The web page is almost totally unreadable since it has a black background with dark gray text. But the worst thing is that there are no links to download the code examples and exercises. While I can read the 3 chapters of updated text in my browser (chrome), the save button doesn't work, I can only print them. There is no link to tell the author any of this, so I have to do it here. I am new to Kindle, but I don't see how I can put this new material back into my Kindle book.")
	// Output:
	// Negative: I just paid good money for this book and went to the web site to download the code examples and exercises. The web page is almost totally unreadable since it has a black background with dark gray text. But the worst thing is that there are no links to download the code examples and exercises. While I can read the 3 chapters of updated text in my browser (chrome), the save button doesn't work, I can only print them. There is no link to tell the author any of this, so I have to do it here. I am new to Kindle, but I don't see how I can put this new material back into my Kindle book.
}

func Test_NBC_4(test *testing.T) {
	// 1. Read/Import training data (DATA) , from my GitHub / Google Docs
	DATA_amazon := bay.GetStruct("../data/train - amazon.csv")
	include := bay.GetInclFt("../data/filter - include.csv")
	exclude := bay.GetExcFt("../data/filter - exclude.csv")

	bay.Print(DATA_amazon, include, exclude, "The book is full of great content and contains the most comprehensive treatment on Go channels and goroutines that I've seen yet. The author also does a phenomenal job with coverage; the table of contents shows this. Virtually every nook and cranny of the Go language is covered. Including the initial founding of Go and its context. (I found this last bit very interesting.) The book is also full of code snippets exemplifying Go idioms and as examples to accomplish common tasks, which is really great considering the infancy of Go.")
	// Output:
	// Positive: The book is full of great content and contains the most comprehensive treatment on Go channels and goroutines that I've seen yet. The author also does a phenomenal job with coverage; the table of contents shows this. Virtually every nook and cranny of the Go language is covered. Including the initial founding of Go and its context. (I found this last bit very interesting.) The book is also full of code snippets exemplifying Go idioms and as examples to accomplish common tasks, which is really great considering the infancy of Go.
}
