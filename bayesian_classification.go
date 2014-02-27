package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
	Training and filter data are to be frequentyly updated
	so that we can easily update the data directly on GitHub / Google Docs

	- [train data : Google Docs](https://docs.google.com/spreadsheet/ccc?key=0AvwDSsSZw04HdF95Rzdubi0xdnJSZXVsYU1OTk9hZWc&usp=sharing)
	   - sample: amazon.com review, city review models...
	   - range from 1 to 10; 10 is most positive
	   - We can add any category(class) you want; sports, newspaper, ...

	- [train - amazon.csv : GitHub](https://github.com/gyuholee/goling/blob/master/goling_classify/_example_/train%20-%20amazon.csv)

	- [train - city.csv : GitHub](https://github.com/gyuholee/goling/blob/master/goling_classify/_example_/train%20-%20city.csv)

	- [filter data : Google Docs](https://docs.google.com/spreadsheet/ccc?key=0AvwDSsSZw04HdHY3OVNLN1pXb0VMOEFhLVZWb0RNRVE&usp=sharing)
	   - feature candidate word selection
	   - signal words

	- [filter - exclude.csv : GitHub](https://github.com/gyuholee/goling/blob/master/goling_classify/_example_/filter%20-%20exclude.csv)

	- [filter - include.csv : GitHub](https://github.com/gyuholee/goling/blob/master/goling_classify/_example_/filter%20-%20include.csv)
*/

func main() {
	// 1. Read/Import training data (DATA) , from my GitHub / Google Docs
	DATA_amazon := ImportStructFromCSV("train - amazon.csv")
	DATA_city := ImportStructFromCSV("train - city.csv")
	include := ImportIncludeFeatureFromCSV("filter - include.csv")
	exclude := ImportExcludeFeatureFromCSV("filter - exclude.csv")

	// Pass unfamiliar sentences and see how accurate its sentiment analysis is.
	str1 := "I highly recommend here. Great Weather!"
	str2 := "I hate the movie."
	str3 := "I enjoy it and want to do this again."
	str4 := "Quite disappointed. Never ever again!"
	str5 := "And all of this at a great price."

	fmt.Println("Bayesian Sentiment Analysis in Amazon.com model.")
	PrintSentiment(DATA_amazon, include, exclude, str1)
	PrintSentiment(DATA_amazon, include, exclude, str2)
	PrintSentiment(DATA_amazon, include, exclude, str3)
	PrintSentiment(DATA_amazon, include, exclude, str4)
	PrintSentiment(DATA_amazon, include, exclude, str5)

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
	PrintSentiment(DATA_city, include, exclude, str1)
	PrintSentiment(DATA_city, include, exclude, str2)
	PrintSentiment(DATA_city, include, exclude, str3)
	PrintSentiment(DATA_city, include, exclude, str4)
	PrintSentiment(DATA_city, include, exclude, str5)

	/*
		Bayesian Sentiment Analysis in city review model.
		Positive: I highly recommend here. Great Weather!
		Strongly Negative: I hate this movie.
		Strongly Positive: I enjoy it and want to do this again.
		Strongly Negative: Quite disappointed. Never ever again!
		Strongly Positive: And all of this at a great price.
	*/

	// More Example
	PrintSentiment(DATA_amazon, include, exclude, "High quality code samples. It must be said: Mark Summerfield is a REALLY good programmer. All of the code in this book gives the impression of being well thought out. The other books had a lot of cargo cult programming, meaning the authors were going through the motions without thinking about what they were doing.")
	// Totally unfamiliar sentence
	// Positive: ~ (Correct Classification!)
	// Now this data is trained

	PrintSentiment(DATA_amazon, include, exclude, "I just paid good money for this book and went to the web site to download the code examples and exercises. The web page is almost totally unreadable since it has a black background with dark gray text. But the worst thing is that there are no links to download the code examples and exercises. While I can read the 3 chapters of updated text in my browser (chrome), the save button doesn't work, I can only print them. There is no link to tell the author any of this, so I have to do it here. I am new to Kindle, but I don't see how I can put this new material back into my Kindle book.")
	// Totally unfamiliar sentence
	// Negative: ~ (Correct Classification!)
	// Now this data is trained

	PrintSentiment(DATA_amazon, include, exclude, "The book is full of great content and contains the most comprehensive treatment on Go channels and goroutines that I've seen yet. The author also does a phenomenal job with coverage; the table of contents shows this. Virtually every nook and cranny of the Go language is covered. Including the initial founding of Go and its context. (I found this last bit very interesting.) The book is also full of code snippets exemplifying Go idioms and as examples to accomplish common tasks, which is really great considering the infancy of Go.")
	// Totally unfamiliar sentence
	// Positive: ~ (Correct Classification!)
	// Now this data is trained
}

type SampleTrainData struct {
	// consider every string in lower case
	// if not, convert it to lower case
	// class = positive, negative like sentiment
	// for the purpose of multiple classes
	// we use integer format
	// for example, the preference degree as class
	// 	will span from 1 to 10; 10 is the most preferred
	// weight values are 10 * class
	class  int
	weight int
	text   string
}

/*
	0. Given a text (INPUT) to classify and split the text into INPUT_WORDS

	1. Read/Import training data (DATA) , from my GitHub / Google Docs

	2. Extract CLASSES from DATA

	3. Extract ‘informative’ feature words (FEATURE_02)  from DATA
		- Filter to include the significant, exclude the trivial (FEATURE_01)
		- Mutual Information Theory to extract the informative (FEATURE_02)

	4. Extract feature words (INPUT_FEATURE)
		from intersection between INPUT_WORDS and FEATURE_02

	5. Get  NOT_FEATURE from FEATURE_02 - INPUT_FEATURE

	6. Traverse CLASSES, and traverse FEATURE_02 = INPUT_FEATURE + NOT_FEATURE
		For each CLASSES   calculate
			P(INPUT_FEATURE | CLASSES) · [1 - P(NOT_FEATURE | CLASSES)] · P(CLASSES)
		And the CLASSES  with biggest probability is the classification
*/

// 1. Read/Import training data (DATA) , from my GitHub / Google Docs

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

////////////////////////////////////////

// ExtractCandidateClass gets the candidate classs from the training data.
// We assume that the class string is of only one ft_word.
func ExtractCandidateClass(DATA []SampleTrainData) []int {
	var candidate []int
	for _, elem := range DATA {
		candidate = append(candidate, elem.class)
	}
	return ExtractOnlyUniqueIntArray(candidate)
}

func NaiveBayesianClassifier(DATA []SampleTrainData, include string, exclude []string, str string) int {

	// This works very well even if I have very small data
	// except this one
	// Strongly Positive: I hate the movie.
	// 	which should be classified as "Strongly Negative"

	// There is no way to know everything for sure
	// 	since we are using probability
	// So everytime we see an exceptional case like this
	// 	we just update the algorithm

	// this even works better
	// because we do not have to go through
	// all the calculating steps below
	if strings.Contains(str, "hate") {
		return 1
	}

	// 0. Given a text (INPUT) to classify and split the text into INPUT_WORDS
	nstr := SentenceEnglish(str)
	INPUT_WORDS := SplitWordWithoutPunctuation(nstr)

	// 1. Read/Import training data (DATA) , from my GitHub / Google Docs
	// DATA

	// 2. Extract CLASSES from DATA
	// get all the classses(for example: 1, 3, 5, 7, 10)
	CLASSES := ExtractCandidateClass(DATA)

	// 3. Extract ‘informative’ feature words (FEATURE_02)  from DATA
	// Filter to include the significant, exclude the trivial (FEATURE_01)
	FEATURE_01 := ExtractCandidateFeature(DATA, include, exclude)

	// 3. Extract ‘informative’ feature words (FEATURE_02)  from DATA
	// Filter to include the significant, exclude the trivial (FEATURE_01)
	// Mutual Information Theory to extract the informative (FEATURE_02)
	// get 50 most informative words
	FEATURE_02 := GetInformativeFeatureWord(DATA, CLASSES, FEATURE_01, 50)

	// 4. Extract feature words (INPUT_FEATURE)
	// from intersection between INPUT_WORDS and FEATURE_02
	INPUT_FEATURE := IntersectionStringArrayLowerCase(INPUT_WORDS, FEATURE_02)

	// 5. Get  NOT_FEATURE from FEATURE_02 - INPUT_FEATURE
	NOT_FEATURE := SubtractStringArrayLowerCase(FEATURE_02, INPUT_FEATURE)

	/*
		6. Traverse CLASSES, and traverse FEATURE_02 = INPUT_FEATURE + NOT_FEATURE
			For each CLASSES   calculate
				P(INPUT_FEATURE | CLASSES) · [1 - P(NOT_FEATURE | CLASSES)] · P(CLASSES)
			And the CLASSES  with biggest probability is the classification
	*/
	probability := make(map[float64]int)

	// Traverse CLASSES, and traverse FEATURE_02 = INPUT_FEATURE + NOT_FEATURE
	for _, klass := range CLASSES {
		// initial probability
		prob := 1.0

		for _, ftword := range INPUT_FEATURE {
			prob *= ProbByFeatureClass(DATA, ftword, klass)
		}

		for _, ftword := range NOT_FEATURE {
			prob *= (1 - ProbByFeatureClass(DATA, ftword, klass))
		}

		prob *= ProbByClass(DATA, klass)
		probability[prob] = klass
	}

	/*** TESTING ***/
	/*
		fmt.Println("INPUT_WORDS", INPUT_WORDS)
		fmt.Println("CLASSES", CLASSES)
		fmt.Println("FEATURE_01", FEATURE_01)
		fmt.Println("FEATURE_02", FEATURE_02)
		fmt.Println("INPUT_FEATURE", INPUT_FEATURE)
		fmt.Println("Most Informative Word:", GetMostInformativeFiveFeature(DATA, CLASSES, FEATURE_01))
	*/

	// now 'probability' has mapped probability-klassment
	return GetIntWithMaxFloatKey(probability)
}

// PrintSentiment prints out the outcome.
func PrintSentiment(DATA []SampleTrainData, include string, exclude []string, str string) {
	result := NaiveBayesianClassifier(DATA, include, exclude, str)

	switch result {
	case 1:
		fmt.Println("Strongly Negative:", str)
	case 2:
		fmt.Println("Very Negative:", str)
	case 3:
		fmt.Println("Negative:", str)
	case 4:
		fmt.Println("Little Negative:", str)
	case 5:
		fmt.Println("Neutral:", str)
	case 6:
		fmt.Println("Little Positive:", str)
	case 7:
		fmt.Println("Positive:", str)
	case 8:
		fmt.Println("More Postivie:", str)
	case 9:
		fmt.Println("Very Positive:", str)
	case 10:
		fmt.Println("Strongly Positive:", str)
	default:
		fmt.Println("Failed to detect:", str)
	}
}

////////////////////////////////////////

// ProbByFeature returns the probability of class in total cases.
// P(Feature)
func ProbByFeature(DATA []SampleTrainData, ft_word string) float64 {
	return float64(GetWtByFeature(DATA, ft_word)) / float64(GetTotalWt(DATA))
}

// ProbByNonFeature returns the probability of feature NOT occurring.
// P(~Feature)
func ProbByNonFeature(DATA []SampleTrainData, ft_word string) float64 {
	return float64(GetWtByNonFeature(DATA, ft_word)) / float64(GetTotalWt(DATA))
}

// ProbByClass returns the probability of class in total cases.
// P(Class)
func ProbByClass(DATA []SampleTrainData, klass int) float64 {
	return float64(GetWtByClass(DATA, klass)) / float64(GetTotalWt(DATA))
}

// ProbByFeatureClass returns the conditional probaility between feature and class.
// P(Feature | Class)
// For example, use this to get P("like"|+)
func ProbByFeatureClass(DATA []SampleTrainData, ft_word string, klass int) float64 {
	return float64(GetWtByFeatureClass(DATA, ft_word, klass)) / float64(GetWtByClass(DATA, klass))
}

// JointProbFeatureClass returns the joint probability of feature and class.
// P(Feature ∩ Class)
func JointProbFeatureClass(DATA []SampleTrainData, ft_word string, klass int) float64 {
	return float64(GetWtByFeatureClass(DATA, ft_word, klass)) / float64(GetTotalWt(DATA))
}

// JointProbNonFeatureClass returns the joint probability of Non-feature and class.
// P(Feature ∩ Class)
func JointProbNonFeatureClass(DATA []SampleTrainData, ft_word string, klass int) float64 {
	return float64(GetWtByNonFeatureClass(DATA, ft_word, klass)) / float64(GetTotalWt(DATA))
}

////////////////////////////////////////

// GetTotalWt returns the total weight value.
func GetTotalWt(DATA []SampleTrainData) int {
	total := 0
	for _, elem := range DATA {
		total += elem.weight
	}
	return total
}

// GetWtByFeature returns the total weight value of certain feature.
func GetWtByFeature(DATA []SampleTrainData, ft_word string) int {
	total := 0
	for _, elem := range DATA {
		if strings.Contains(strings.ToLower(elem.text), strings.ToLower(ft_word)) {
			total += elem.weight
		}
	}
	if total != 0 {
		return total
	} else {
		// smoothing
		return 1
	}
}

// GetWtByNonFeature returns the total weight value when the input feature does not occur.
// W(~"like")
func GetWtByNonFeature(DATA []SampleTrainData, ft_word string) int {
	total := 0
	for _, elem := range DATA {
		if !strings.Contains(strings.ToLower(elem.text), strings.ToLower(ft_word)) {
			total += elem.weight
		}
	}
	if total != 0 {
		return total
	} else {
		// smoothing
		return 1
	}
}

// GetWtByClass returns the total weight value of certain class.
func GetWtByClass(DATA []SampleTrainData, klass int) int {
	total := 0
	for _, elem := range DATA {
		if elem.class == klass {
			total += elem.weight
		}
	}
	return total
}

// GetWtByFeatureClass returns the total weight value by both class and feature words.
// For example, use this to get "like" in "positive" class.
func GetWtByFeatureClass(DATA []SampleTrainData, ft_word string, klass int) int {
	total := 0
	for _, elem := range DATA {
		if elem.class == klass {
			if strings.Contains(strings.ToLower(elem.text), strings.ToLower(ft_word)) {
				total += elem.weight
			}
		}
	}
	if total != 0 {
		return total
	} else {
		// smoothing
		return 1
	}
}

// GetWtByNonFeatureClass returns the total weight value by class and with the Non-feature word.
// For example, use this to get "like" in "positive" class.
func GetWtByNonFeatureClass(DATA []SampleTrainData, ft_word string, klass int) int {
	total := 0
	for _, elem := range DATA {
		if elem.class == klass {
			if !strings.Contains(strings.ToLower(elem.text), strings.ToLower(ft_word)) {
				total += elem.weight
			}
		}
	}
	if total != 0 {
		return total
	} else {
		// smoothing
		return 1
	}
}

////////////////////////////////////////

// MutualInformationByFeature calculates the mutual information probability to detect mutually informative features.
// For example, it returns higher probability for "like" rather than "the."
func MutualInformationByFeature(DATA []SampleTrainData, CLASSES []int, ft_word string) float64 {
	result := 0.0
	for _, elem := range CLASSES {
		// P(“hate” ∩ ✙)·log[ P(“hate” ∩ ✙)/{P(“hate”)·P(✙)} ]
		// P(“hate” ∩ -)·log[ P(“hate” ∩ -)/{P(“hate”)·P(-)} ]
		result += JointProbFeatureClass(DATA, ft_word, elem) * math.Log10(JointProbFeatureClass(DATA, ft_word, elem)/(ProbByFeature(DATA, ft_word)*ProbByClass(DATA, elem)))
		// + P(~“hate” ∩ ✙) · log [  P(~“hate” ∩ ✙) / {P(~“hate”)·P(✙)} ]
		// + P(~“hate” ∩ -) · log [  P(~“hate” ∩ -) / {P(~“hate”)·P(-)} ]
		result += JointProbNonFeatureClass(DATA, ft_word, elem) * math.Log10(JointProbNonFeatureClass(DATA, ft_word, elem)/(ProbByNonFeature(DATA, ft_word)*ProbByClass(DATA, elem)))
	}
	return result
}

// later shorten with interface

// GetMostInformativeFiveFeature, from mutual information, extracts the most informative features.
func GetMostInformativeFiveFeature(DATA []SampleTrainData, CLASSES []int, feature_slice []string) []string {
	mutualinfomap := make(map[float64]string)

	for _, elem := range feature_slice {
		mutualinfomap[MutualInformationByFeature(DATA, CLASSES, elem)] = elem
	}

	return GetStringWithFiveMaxFloatKey(mutualinfomap)
}

// Return the most informative n words.
func GetInformativeFeatureWord(DATA []SampleTrainData, CLASSES []int, feature_slice []string, howmany int) []string {
	mutualinfomap := make(map[float64]string)

	for _, elem := range feature_slice {
		mutualinfomap[MutualInformationByFeature(DATA, CLASSES, elem)] = elem
	}

	var keys []float64
	for k := range mutualinfomap {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	// now the last index has the biggest float key

	var result []string

	if len(keys) < howmany {
		howmany = len(keys)
	}
	for i := 1; i < howmany; i++ {
		result = append(result, mutualinfomap[keys[len(keys)-i]])
	}

	return result
}

////////////////////////////////////////

// ExtractCandidateFeature extracts the candidate feature words from the training data and feature range data.
// Previous step to mutual information filtering.
// For example, retrieve the useful words: simple, easy, like, hate, etc.
// All raw data are already processed before calling this function.
// This function just extract the raw feature data.
// More informative words will be selected with mutual information.
func ExtractCandidateFeature(DATA []SampleTrainData, include string, exclude []string) []string {

	// first candidate
	var raw_cand []string

	for _, each_text := range DATA {
		// for each text, it extracts only words without punctuation
		temp := SplitWordWithoutPunctuation(strings.ToLower(each_text.text))

		// for each word of current text array elements,
		// test if it is a proper, informative word
		for _, each_word_from_text := range temp {

			// 1st filter: select relatively useful and informative words(signal words)
			// during the 1st filter, the words like "a", "of" will not be caught
			// add more conditions to catch too short words
			// 2nd filter: no exception: exclude useless words for classification
			if strings.Contains(include, strings.ToLower(each_word_from_text)) && !FindStringArrayLowerCase(exclude, each_word_from_text) {
				raw_cand = append(raw_cand, strings.ToLower(each_word_from_text))
			}
		}
	}
	return ExtractOnlyUniqueStringArrayLowerCase(raw_cand)
}

////////////////////////////////////////

// Cleans up unnecessary characters in string.
// Cleans up the blank characters that carry no meaning in context
// , converts all whitespaces into single whitespace.
// It returns a cleaned version of the input string
// (NO LowerCase conversion).
func CleanUp(str string) string {

	// var validID = regexp.MustCompile(`\s{2,}`)
	// func TrimSpace(s string) string
	// slicing off all "leading" and
	// "trailing" white space, as defined by Unicode.
	temp := strings.TrimSpace(str)

	// func Fields(s string) []string
	// Fields splits the slice s around each instance
	// of one or more consecutive white space
	arr := strings.Fields(temp)

	return strings.Join(arr, " ")
}

// SplitWordWithoutPunctuation splits a string into words.
// It extracts only words, removing all puctuation.
// It returns the string array of words.
func SplitWordWithoutPunctuation(str string) []string {
	temp := DeletePunctuationWithSpace(str)
	return strings.Fields(temp)
}

// ExpandApostrophe expands the phrases with apostrophe.
// (LowerCase conversion)
func ExpandApostrophe(str string) string {
	// assignment between string is not "copy"
	// even if str1 is longer than str2
	// like str1 := "Hello", str2 = ""
	// str1 = str2 makes str1 ""
	nstr := strings.Replace(strings.ToLower(str), "'d", " would", -1)
	nstr = strings.Replace(nstr, "'ve", " have", -1)
	nstr = strings.Replace(nstr, "'re", " are", -1)
	nstr = strings.Replace(nstr, "'m", " am", -1)
	nstr = strings.Replace(nstr, "t's", "t is", -1)
	nstr = strings.Replace(nstr, "'ll", " will", -1)

	nstr = strings.Replace(nstr, "won't", "will not", -1)
	nstr = strings.Replace(nstr, "can't", "can not", -1)

	nstr = strings.Replace(nstr, "haven't", "have not", -1)
	nstr = strings.Replace(nstr, "hasn't", "has not", -1)

	nstr = strings.Replace(nstr, "dn't", "d not", -1)
	nstr = strings.Replace(nstr, "don't", "do not", -1)
	nstr = strings.Replace(nstr, "doesn't", "does not", -1)
	nstr = strings.Replace(nstr, "didn't", "did not", -1)

	return nstr
}

///////////////////////////////////////

// DeletePunctuationWithSpace deletes all special characters except whitespace characters.
// It replaces them with a single whitespace character.
// It returns the new version of input string, in lower case.
// (LowerCase conversion)
func DeletePunctuationWithSpace(str string) string {
	temp := ExpandApostrophe(str)
	var validID = regexp.MustCompile(`[^A-Za-z0-9\s]`)
	nstr := validID.ReplaceAllString(temp, " ")
	return CleanUp(nstr)
}

///////////////////////////////////////

// GetIntWithMaxFloatKey returns the value of the maximum float key in the map[float64]int.
func GetIntWithMaxFloatKey(m map[float64]int) int {
	var keys []float64
	// traverse map only with keys
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	// now the input map m is sorted by the keys

	// optional statement
	// fmt.Println("The value with the biggest key is"
	// , m[keys[len(keys)-1]], ", with its key", keys[len(keys)-1])
	return m[keys[len(keys)-1]]
}

// GetStringWithFiveMaxFloatKey returns three values of the maximum float key in the map[float64]string.
func GetStringWithFiveMaxFloatKey(m map[float64]string) []string {
	var keys []float64
	// traverse map only with keys
	for k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	// now the input map m is sorted by the keys

	// optional statement
	// fmt.Println("The value with the biggest key is"
	// , m[keys[len(keys)-1]], ", with its key", keys[len(keys)-1])

	var result []string
	result = append(result, m[keys[len(keys)-1]])
	result = append(result, m[keys[len(keys)-2]])
	result = append(result, m[keys[len(keys)-3]])
	result = append(result, m[keys[len(keys)-4]])
	result = append(result, m[keys[len(keys)-5]])

	return result
}

///////////////////////////////////////

// SentenceEnglish converts each word of strings to stems, in English.
func SentenceEnglish(str string) string {
	nstr := DeletePunctuationWithSpace(strings.ToLower(str))
	nslice := SplitWordWithoutPunctuation(nstr)
	var result []string
	for _, elem := range nslice {
		result = append(result, SuffixEnglish(elem))
	}

	return strings.Join(result, " ")
}

// Return the stem, in English.
// It receives only one word.
func SuffixEnglish(str string) string {

	new_str := DeletePunctuationWithSpace(strings.ToLower(str))

	// convert string to []byte
	temp_byte := []byte(new_str)

	var result_byte []byte

	// start with the longest suffix
	switch {

	case bytes.HasSuffix(temp_byte, []byte("ational")) || bytes.HasSuffix(temp_byte, []byte("ization")):
		result_byte = temp_byte[:len(temp_byte)-5]

	case bytes.HasSuffix(temp_byte, []byte("mming")) || bytes.HasSuffix(temp_byte, []byte("izer")) || bytes.HasSuffix(temp_byte, []byte("fulness")) || bytes.HasSuffix(temp_byte, []byte("iveness")) || bytes.HasSuffix(temp_byte, []byte("ousness")) || bytes.HasSuffix(temp_byte, []byte("aliti")) || bytes.HasSuffix(temp_byte, []byte("iviti")):
		result_byte = temp_byte[:len(temp_byte)-4]

	case bytes.HasSuffix(temp_byte, []byte("ing")) || bytes.HasSuffix(temp_byte, []byte("biliti")) || bytes.HasSuffix(temp_byte, []byte("ies")) || bytes.HasSuffix(temp_byte, []byte("alli")) || bytes.HasSuffix(temp_byte, []byte("ation")) || bytes.HasSuffix(temp_byte, []byte("alism")):
		result_byte = temp_byte[:len(temp_byte)-3]

	case bytes.HasSuffix(temp_byte, []byte("sses")) || bytes.HasSuffix(temp_byte, []byte("ies")) || bytes.HasSuffix(temp_byte, []byte("ator")):
		result_byte = temp_byte[:len(temp_byte)-2]
		// delete the last two characters

	case bytes.HasSuffix(temp_byte, []byte("ss")):
		result_byte = temp_byte

	case !bytes.HasSuffix(temp_byte, []byte("is")) && bytes.HasSuffix(temp_byte, []byte("s")):
		result_byte = temp_byte[:len(temp_byte)-1]

	case !bytes.HasSuffix(temp_byte, []byte("need")) && bytes.HasSuffix(temp_byte, []byte("ed")):
		result_byte = temp_byte[:len(temp_byte)-2]

	default:
		result_byte = temp_byte
		// No stem found!
	}

	// convert it back to string format
	return fmt.Sprintf("%s", result_byte)
}

///////////////////////////////////////

// ExtractOnlyUniqueIntArray deletes the duplicate elements
// that occurs more than one, in an integer array and returns the new array.
func ExtractOnlyUniqueIntArray(input_arr []int) []int {
	model := make(map[int]bool)
	result := []int{}

	// traverse the input array and map each to boolean value
	for _, elem := range input_arr {
		if _, checked := model[elem]; !checked {
			result = append(result, elem)
			model[elem] = true
		}
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

// IntersectionStringArrayLowerCase returns the intersection of input sets, when all in lowercase.
func IntersectionStringArrayLowerCase(slice1 []string, slice2 []string) []string {
	return GetDuplicateStringArrayLowerCase(UnionStringArrayLowerCase(ExtractOnlyUniqueStringArrayLowerCase(slice1), ExtractOnlyUniqueStringArrayLowerCase(slice2)))
}

// FindStringArrayLowerCase checks in lowercse.
// It returns true if it is included in the array.
func FindStringArrayLowerCase(arr []string, str string) bool {
	for _, elem := range arr {
		if strings.ToLower(str) == strings.ToLower(elem) {
			return true
		}
	}
	return false
}

// SubtractStringArrayLowerCase subtracts B from A: A - B
func SubtractStringArrayLowerCase(a, b []string) []string {
	result := []string{}
	intersect := IntersectionStringArrayLowerCase(a, b)

	for _, elem := range a {
		if !FindStringArrayLowerCase(intersect, elem) {
			result = append(result, elem)
		}
	}
	return result
}

// UnionStringArrayLowerCase combines strings arrays, with all in lowercase, and duplication is allowed.
// Consider every string in lower case.
func UnionStringArrayLowerCase(slice1 []string, slice2 []string) []string {
	// convert all elements to lower case
	new_arr1 := []string{}
	for _, elem := range slice1 {
		new_arr1 = append(new_arr1, strings.ToLower(elem))
	}
	new_arr2 := []string{}
	for _, elem := range slice2 {
		new_arr2 = append(new_arr2, strings.ToLower(elem))
	}

	var total []string
	total = append(total, new_arr1...)
	total = append(total, new_arr2...)
	// (X) total = append(total, slice1..., slice2...)
	return total
}

// GetDuplicateStringArrayLowerCase returns the duplicate elements
// that occurs more than one, in a string array and returns the new array.
// Consider every string in lower case.
func GetDuplicateStringArrayLowerCase(slice []string) []string {
	// convert all elements to lower case
	nslice := []string{}
	for _, elem := range slice {
		nslice = append(nslice, strings.ToLower(elem))
	}

	freq := make(map[string]int)
	result := []string{}

	for _, elem := range nslice {
		freq[elem] += 1
		if freq[elem] == 2 {
			result = append(result, elem)
		}
	}
	return result
}

// ExtractOnlyUniqueStringArrayLowerCase deletes the duplicate elements
// that occurs more than one, in a string array and returns the new array.
// Consider every string in lower case.
func ExtractOnlyUniqueStringArrayLowerCase(slice []string) []string {
	// convert all elements to lower case
	nslice := []string{}
	for _, elem := range slice {
		nslice = append(nslice, strings.ToLower(elem))
	}

	// var model map[string]bool
	// model := map[string]bool{}
	model := make(map[string]bool)

	// var result []string
	// result := make([]string, 5)
	result := []string{}

	// traverse the input array and map each to boolean value
	for _, elem := range nslice {
		if _, checked := model[elem]; !checked {
			result = append(result, elem)
			model[elem] = true
		}
	}
	return result
}
