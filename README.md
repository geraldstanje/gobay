gobay [![Build Status](https://travis-ci.org/gyuho/gobay.png?branch=master)](https://travis-ci.org/gyuho/gobay) [![GoDoc](https://godoc.org/github.com/gyuho/gobay?status.png)](http://godoc.org/github.com/gyuho/gobay) [![Project Stats](http://www.ohloh.net/p/714281/widgets/project_thin_badge.gif)](http://www.ohloh.net/p/714281)
==========

Package gobay is a small package for sentiment analysis using Bayesian probability.


YouTube Clips by me
==========
<a href="http://www.youtube.com/watch?v=dctzCcYt4AM" target="_blank"><img src="http://img.youtube.com/vi/dctzCcYt4AM/0.jpg"></a>
<ul>
	<li><a href="https://www.youtube.com/watch?v=dctzCcYt4AM&list=PLT6aABhFfintOGKWVWz9qMxC3qZZdHQRD&index=1" target="_blank">Bayesian Classification</li>
	<li><a href="https://www.youtube.com/watch?v=927YDZH_MLo&list=PLT6aABhFfintOGKWVWz9qMxC3qZZdHQRD" target="_blank">String Similarity, Cosine Similarity, Levenshtein Distance</li>
	<li><a href="https://www.youtube.com/watch?v=3qHx1VCcobY&list=PLT6aABhFfintOGKWVWz9qMxC3qZZdHQRD" target="_blank">Spell Check</li>
</ul>

[↑ top](https://github.com/gyuho/gobay#gobay---)


Getting Started
==========
- [godoc.org](http://godoc.org/github.com/gyuho/gobay)
- [gowalker.org](http://gowalker.org/github.com/gyuho/gobay#_index)

```go
// to install, in the command line
mkdir $HOME/go
export GOPATH=$HOME/go
go get github.com/gyuho/gobay

// to include, in the code
import "github.com/gyuho/gobay"

// to call the function, in the code
[package_name].[function]

// to run, or go install
go run [path/filename]
```
[↑ top](https://github.com/gyuho/gobay#gobay---)


gobay Package Hierarchy
==========
```go
bay/		# Naive Bayesian Classifier
data/		# Training Data
read/		# Import Training Data
slm/		# Slice, Map Functions
```
[↑ top](https://github.com/gyuho/gobay#gobay---)


Example
==========
```go
package gobay
```
[↑ top](https://github.com/gyuho/gobay#gobay---)


Training Data
==========
Training and filter data are to be frequentyly updated, directly from GitHub / Google Docs.

- <a href="https://docs.google.com/spreadsheet/ccc?key=0AvwDSsSZw04HdF95Rzdubi0xdnJSZXVsYU1OTk9hZWc&usp=sharing" target="_blank">train data : Google Docs</a>
	- sample: amazon.com review, city review models...
	- range from 1 to 10; 10 is most positive
	- We can add any category(class) you want; sports, newspaper, ...

- <a href="https://github.com/gyuholee/gobay/blob/master/data/train%20-%20amazon.csv" target="_blank">train - amazon.csv : GitHub</a>

- <a href="https://github.com/gyuholee/gobay/blob/master/data/train%20-%20city.csv" target="_blank">train - city.csv : GitHub</a>

- <a href="https://docs.google.com/spreadsheet/ccc?key=0AvwDSsSZw04HdHY3OVNLN1pXb0VMOEFhLVZWb0RNRVE&usp=sharing" target="_blank">filter data : Google Docs</a>
	- feature candidate word selection
	- signal words

- <a href="https://github.com/gyuholee/gobay/blob/master/data/filter%20-%20exclude.csv" target="_blank">filter - exclude.csv : GitHub</a>

- <a href="https://github.com/gyuholee/gobay/blob/master/data/filter%20-%20include.csv" target="_blank">filter - include.csv : GitHub</a>


[↑ top](https://github.com/gyuho/gobay#gobay---)

To-Do-List
==========
- Update Bayesian algorithms for some exceptional cases

[↑ top](https://github.com/gyuho/gobay#gobay---)
