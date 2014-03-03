package read

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// GetStruct imports data from a csv file and construct the structure.
func GetStruct(filename string) []SampleTrainData {
	output := CSV(filename)
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

// GetInclFt imports "include" feature candidate range data from a csv file.
// Possibly big file, so use strings.Contains method should be faster.
func GetInclFt(filename string) string {
	output := CSV(filename)
	var include_str string

	// row traverse
	// count excluding the header row
	for i := 1; i < len(output); i++ {
		include_str += "," + output[i][0]
	}
	return include_str
}

// GetExcFt imports "exclude" feature candidate range data from a csv file.
// Relatively small amount of data.
// Just to be used with linear search.
func GetExcFt(filename string) []string {
	output := CSV(filename)
	var exclude_arr []string

	// row traverse
	// count excluding the header row
	for i := 1; i < len(output); i++ {
		exclude_arr = append(exclude_arr, output[i][0])
	}

	return exclude_arr
}

// CSV reads data from a csv file.
// [][]string:
// the first [] is row, the second [] is column.
// len(output) would be the number of total rows.
// Use the following line to traverse
// by all rows and only the first column.
// for i := 0; i < len(output); i++
// 	output[i][0]
func CSV(filename string) [][]string {
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
