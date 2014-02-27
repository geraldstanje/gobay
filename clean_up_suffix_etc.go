package main

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice1 := []string{"hello", "Hello", "Like", "world", "Good", "time", "GO"}
	slice2 := []string{"hello", "Hello"}
	fmt.Println(SubtractStringArrayLowerCase(slice1, slice2))
	// [Like world Good time GO]

	fmt.Println(IntersectionStringArrayLowerCase(slice1, slice2))
	// [hello]
}

// CleanUp cleans up unnecessary characters in string.
// It cleans up the blank characters that carry no meaning in context
// , converts all whitespaces into single whitespace.
// String is immutable, which means the original string would not change.
func CleanUp(str string) string {

	// validID := regexp.MustCompile(`\s{2,}`)
	// func TrimSpace(s string) string
	// slicing off all "leading" and
	// "trailing" white space, as defined by Unicode.
	str = strings.TrimSpace(str)

	// func Fields(s string) []string
	// Fields splits the slice s around each instance
	// of "one or more consecutive white space"
	slice := strings.Fields(str)

	// now join them with a single white space character
	return strings.Join(slice, " ")
}

// GetWords extracts only words, removing all puctuation.
// It returns an slice of words.
func GetWords(str string) []string {
	str = ExpandApostrophe(str)
	str = ReplaceNonAlnumWithSpace(str)
	return strings.Fields(str)
}

// ExpandApostrophe expands the apostrophe phrases.
// And convert them to lower case letters.
func ExpandApostrophe(str string) string {
	// assignment between string is not "copy"
	// even if str1 is longer than str2
	// like str1 := "Hello", str2 = ""
	// str1 = str2 makes str1 ""
	str = strings.Replace(strings.ToLower(str), "'d", " would", -1)

	// If n < 0, there is no limit on the number of replacements.
	str = strings.Replace(str, "'ve", " have", -1)
	str = strings.Replace(str, "'re", " are", -1)
	str = strings.Replace(str, "'m", " am", -1)
	str = strings.Replace(str, "t's", "t is", -1)
	str = strings.Replace(str, "'ll", " will", -1)

	str = strings.Replace(str, "won't", "will not", -1)
	str = strings.Replace(str, "can't", "can not", -1)

	str = strings.Replace(str, "haven't", "have not", -1)
	str = strings.Replace(str, "hasn't", "has not", -1)

	str = strings.Replace(str, "dn't", "d not", -1)
	str = strings.Replace(str, "don't", "do not", -1)
	str = strings.Replace(str, "doesn't", "does not", -1)
	str = strings.Replace(str, "didn't", "did not", -1)

	return str
}

///////////////////////////////////////

// DeletePunctuationWithSpace deletes all special
// characters except whitespace characters.
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

// GetIntWithMaxFloatKey returns the value of
// the maximum float key in the map[float64]int.
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

// GetStringWithFiveMaxFloatKey returns three values
// of the maximum float key in the map[float64]string.
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
	nslice := GetWords(nstr)
	var result []string
	for _, elem := range nslice {
		result = append(result, SuffixEnglish(elem))
	}

	return strings.Join(result, " ")
}

// SuffixEnglish return the stem, in English.
// It receives only one word.
func SuffixEnglish(str string) string {

	ns := DeletePunctuationWithSpace(strings.ToLower(str))

	// convert string to []byte
	tm := []byte(ns)

	var result_byte []byte

	// start with the longest suffix
	switch {

	case bytes.HasSuffix(tm, []byte("ational")) || bytes.HasSuffix(tm, []byte("ization")):
		result_byte = tm[:len(tm)-5]

	case bytes.HasSuffix(tm, []byte("mming")) || bytes.HasSuffix(tm, []byte("izer")) || bytes.HasSuffix(tm, []byte("fulness")) || bytes.HasSuffix(tm, []byte("iveness")) || bytes.HasSuffix(tm, []byte("ousness")) || bytes.HasSuffix(tm, []byte("aliti")) || bytes.HasSuffix(tm, []byte("iviti")):
		result_byte = tm[:len(tm)-4]

	case bytes.HasSuffix(tm, []byte("ing")) || bytes.HasSuffix(tm, []byte("biliti")) || bytes.HasSuffix(tm, []byte("ies")) || bytes.HasSuffix(tm, []byte("alli")) || bytes.HasSuffix(tm, []byte("ation")) || bytes.HasSuffix(tm, []byte("alism")):
		result_byte = tm[:len(tm)-3]

	case bytes.HasSuffix(tm, []byte("sses")) || bytes.HasSuffix(tm, []byte("ies")) || bytes.HasSuffix(tm, []byte("ator")):
		result_byte = tm[:len(tm)-2]
		// delete the last two characters

	case bytes.HasSuffix(tm, []byte("ss")):
		result_byte = tm

	case !bytes.HasSuffix(tm, []byte("is")) && bytes.HasSuffix(tm, []byte("s")):
		result_byte = tm[:len(tm)-1]

	case !bytes.HasSuffix(tm, []byte("need")) && bytes.HasSuffix(tm, []byte("ed")):
		result_byte = tm[:len(tm)-2]

	default:
		result_byte = tm
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

// IntersectionStringArrayLowerCase returns the
// intersection of input sets, when all in lowercase.
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

// UnionStringArrayLowerCase combines strings arrays
// , with all in lowercase, and duplication is allowed.
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
