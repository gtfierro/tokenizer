package tokenizer

import (
	"bufio"
	"fmt"
	"github.com/agonopol/go-stem/stemmer"
	"html"
	"os"
	"strconv"
	"strings"
)

var replacer = strings.NewReplacer(".", "", ",", "", "!", "", "?", "", "||", "", "(", "", ")", "")

var Dict = make(map[string]int) // maps tokens to an index
var Dictfile = "dict.csv"
var Matrixfile = "matrix.csv"

/**
  Takes as input a filename containing the full text for a patent on each line. Returns
  a string slice where each entry is the full text of a patent (order is maintained)
  which has had the following transformations applied to it:
  * unescaped HTML sequences
  * lowercase
  * removed all .,!?||()
*/
func Read_file(filename string) []string {
	results := []string{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = html.UnescapeString(line)
		line = strings.ToLower(line)
		line = replacer.Replace(line)
		results = append(results, line)
	}
	fmt.Println("Finished reading input file", filename)
	return results
}

/**
  Takes as input a space-delimited string and returns a slice of all the
  tokens in that string.

  If the `stem` flag is True, applies the Porter stemming algorithm
  to each token
*/
func tokenize(instring string, stem bool) []string {
	tokens := strings.Split(instring, " ")
	if stem {
		res := []string{}
		for _, t := range tokens {
			stem := stemmer.Stem([]byte(t))
			res = append(res, string(stem))
		}
		return res
	}
	return tokens
}

/**
  TODO: fold this functionality into reading the file

  Given the output of Read_file, populates Dict, a map[string]int.
  Iterates through each of the found lines, tokenizes the line,
  and adds tokens to the Dict, which maintains a mapping of a token
  to its index
*/
func CreateDict(lines []string) {
	fmt.Println("Creating token dictionary")
	index := 0
	for _, line := range lines {
		tokens := tokenize(line, false)
		for _, token := range tokens {
			if Dict[token] == 0 {
				Dict[token] = index
				index += 1
			}
		}
	}
	fmt.Println("Finished creating token dictionary with", len(Dict), "items")
}
