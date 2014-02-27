package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	dslice_amazon := ImportStructFromCSV("train - amazon.csv")
	dslice_city := ImportStructFromCSV("train - city.csv")
	include_str := ImportIncludeFeatureFromCSV("filter - include.csv")
	exclude_arr := ImportExcludeFeatureFromCSV("filter - exclude.csv")

	fmt.Println(dslice_amazon)
	println()
	/*
	   [{1 100 I am tired of this movie.} {1 100 I hate this movie.} {1 100 Not really happy about this at all.} {1 100 Terrible buying experience.} {1 100 The write-up was very misleading.} {1 100 This is by far the worst I've ever seen.} {1 100 Throughly disappointed.} {1 100 total scam} {1 100 Do not buy it!} {1 100 I also thought it was too verbose and repetitive.} {1 100 The web page is almost totally unreadable since it has a black background with dark gray text.} {1 100 But the worst thing is that there are no links to download the code examples and exercises.} {3 300 I will not pay a dollar more.} {3 300 Quite disappointed.} {3 300 The text provides few examples and relies far too much on its problems as a teaching aid.} {3 300 It is extra unimportant features.} {3 300 Quite disappointing.} {3 300 I needed an external source to help me deal with cases not covered in the text.} {3 300 This book obviously became a cult and spreads a very dangerous trend. } {3 300 They won't fail the interview but they are not good at solving the simplest real-world problems} {3 300 My biggest complaint is the organization.} {3 300 And it's hard to tell who the target audience is.} {3 300 However, you should note that the book does not actually contain the example programs. } {3 300 While I can read the 3 chapters of updated text in my browser (chrome), the save button doesn't work, I can only print them.} {3 300 There is no link to tell the author any of this, so I have to do it here.} {5 500 The only negative thing about this book is the lack of solutions to exercises.} {5 500 Presentation Could Be Better} {7 700 I have very few complaints, and this text has been one of my favorites throughout college.} {7 700 I needed an external source to help me deal with cases not covered in the text.} {7 700 I feel very good about its story.} {7 700 I love the way that the scenes are identified as new or extended.} {7 700 A must book for the people who want to make it into the BIG software companies.} {7 700 the paper quality is quite good and the words are very clear, hope this book can assist me to master the knowledge and find a good job!} {7 700 very useful especially for whom never really has chance to know what is going about get a job} {7 700 Dr. Aziz uses a writing style that is succinct and to the point.} {7 700 Compared to that book, I found Elements of Programming Interviews to be better preparation for the questions I was asked at the on-site interviews.} {7 700 It must be said: Mark Summerfield is a REALLY good programmer} {7 700 All of the code in this book gives the impression of being well thought out} {7 700 I found this last bit very interesting.} {7 700 The book is also full of code snippets exemplifying Go idioms and as examples to accomplish common tasks, which is really great considering the infancy of Go.} {7 700 Virtually every nook and cranny of the Go language is covered.} {10 1000 So happy to have this in my movie collection.} {10 1000 A great edition, which will grace any collection! } {10 1000 I highly recommended this movie trilogy.} {10 1000 The quality of sound and image perfectly.} {10 1000 There is no word to describe how much I love this move.} {10 1000 This is an absolute masterpiece.} {10 1000 This is the best epic movie I have ever watched.} {10 1000 I could not be more happy with it.} {10 1000 I have no doubt this is an excellent introduction, reference.} {10 1000 The author also does a phenomenal job with coverage; the table of contents shows this.} {10 1000 The book is full of great content and contains the most comprehensive treatment on Go channels and goroutines that I've seen yet.}]
	*/

	fmt.Println(dslice_city)
	println()
	/*
	   [{1 100 I am tired of this city.} {1 100 I don't like this city} {1 100 I hate this city.} {1 100 I won't come back!} {1 100 Quite disappointed.} {1 100 Quite disappointing.} {1 100 The weather is terrible.} {3 300 I can't deal with this town anymore.} {3 300 Venice is notoriously expensive. It's too pricey.} {5 500 It's too cold!} {5 500 The weather is okay.} {5 500 Not much to do.} {5 500 It is an okay city.} {5 500 I visited the city.} {7 700 Dining out in Venice is expensive so it comes as a surprise to discover that drinking is incredibly cheap, which may explain why locals can often be spotted with a glass in hand from 10am.} {7 700 Everything is cheap here.} {7 700 I feel very good about its food and atmosphere.} {7 700 I highly recommend this city.} {7 700 I like the weather here.} {7 700 In the quiet village of Binfield, this 17th-century farmhouse has a gorgeous garden and two charmingly converted buildings. } {7 700 The price is very cheap.} {7 700 Well, not arrived, exactly. The Idle Rocks was sitting on St Mawes harbour, hogging the best views, long before I first came in 2001.} {7 700 Your money goes much further if you can escape the centre and head into one of the city's residential neighbourhoods – which is fun to do without a map; more than any other city, there's a surprise on every corner in Venice.} {7 700 I had a great time here.} {7 700 It was exciting.} {10 1000 Definitely want to visit again.} {10 1000 I love the weather here.} {10 1000 One of the best cities I've ever been.} {10 1000 The location is very accessible.} {10 1000 This is an amazing place!} {10 1000 This is a great city.} {10 1000 This is the best city.} {10 1000 This is an awesome city. I am having an amazing time.} {10 1000 It was an awesome experience.}]
	*/

	fmt.Println(include_str)
	println()
	/*
	   ,able,absolute,amazing,appropriate,awesome,bad,beautiful,benefit,best,better,blowing,cheap,classic,clear,compact,compare,daunting,decent,definitely,disappoint,disappointed,disappointing,enjoy,epic,error,even,ever,every,excellent,exciting,extra,far,favorite,feel,genuine,good,grand,great,greatest,happy,harmful,hate,here,highly,honest,illogical,inexpensive,interested,like,lot,lots,love,lovely,loving,main,masterpiece,mind,mindblowing,misleading,more,most,much,must,never,no,not,obvious,perfectly,point,pretty,quality,quite,really,recommended,respect,scam,scary,simple,simply,stars,strong,succinct,suggest,sure,terrible,there,thoroughly,tired,total,unimportant,useful,very,visit,well,winning,worse,worst,worth,worthwhile
	*/

	fmt.Println(exclude_arr)
	/*
	   [a all am an and are as at be because for from has have hi i in into is it may might movies of off on or out pages scenes since that the they this to too us was way were with you]
	*/
}

type SampleTrainData struct {
	// consider every string in lower case
	// if not, convert it to lower case
	// class = positive, negative like sentiment
	// for the purpose of multiple classes
	// we use integer format
	// for example, the preference degree as class
	// will span from 1 to 10; 10 is the most preferred
	class  int
	weight int
	text   string
}

// ImportStructFromCSV imports data from a csv file and construct the structure.
func ImportStructFromCSV(filename string) []SampleTrainData {
	output := ImportCSV(filename)
	var struct_arr SampleTrainData
	var result []SampleTrainData

	// row traverse
	// count excluding the header row
	for i := 1; i < len(output); i++ {
		struct_arr.class = StrToInt(output[i][0])
		struct_arr.weight = StrToInt(output[i][0]) * 100
		struct_arr.text = output[i][1]
		result = append(result, struct_arr)
	}
	return result
}

// StrToInt converts string to integer.
func StrToInt(str string) int {
	i, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err.Error())
		panic("Fail")
	}

	return i
}

// ImportIncludeFeatureFromCSV imports "include" feature candidate range data from a csv file.
// Possibly big file, so use strings.Contains method should be faster.
func ImportIncludeFeatureFromCSV(filename string) string {
	output := ImportCSV(filename)
	var include_str string

	// row traverse
	// count excluding the header row
	for i := 1; i < len(output); i++ {
		include_str += "," + output[i][0]
	}
	return include_str
}

// ImportExcludeFeatureFromCSV imports "exclude" feature candidate range data from a csv file.
// Relatively small amount of data.
// Just to be used with linear search.
func ImportExcludeFeatureFromCSV(filename string) []string {
	output := ImportCSV(filename)
	var exclude_arr []string

	// row traverse
	// count excluding the header row
	for i := 1; i < len(output); i++ {
		exclude_arr = append(exclude_arr, output[i][0])
	}

	return exclude_arr
}

// ImportCSV reads data from a csv file.
// [][]string:
// the first [] is row, the second [] is column.
// len(output) would be the number of total rows.
// Use the following line to traverse
// by all rows and only the first column.
// for i := 0; i < len(output); i++
// 	output[i][0]
func ImportCSV(filename string) [][]string {
	// func Open(name string) (file *File, err error)
	file, err := os.Open(filename)
	if err != nil {
		// fmt.Println(err.Error())
		fmt.Println("Error:", err)
		return nil
	}

	// func (f *File) Close() error
	defer file.Close()

	// func NewReader(r io.Reader) *Reader
	// it can read csv or txt file
	reader := csv.NewReader(file)

	reader.TrailingComma = true
	reader.TrimLeadingSpace = true
	// reader.LazyQuotes = true

	for {
		// func (r *Reader) ReadAll() (records [][]string, err error)
		data, read_err := reader.ReadAll()

		if read_err == io.EOF {
			break
		} else if read_err != nil {
			// fmt.Println(err.Error())
			fmt.Println("Error:", read_err)
			return nil
		}
		return data
	}
	return nil
}
