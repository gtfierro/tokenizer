package tokenizer

import (
	"bufio"
	"fmt"
	"github.com/agonopol/go-stem/stemmer"
	"html"
	"os"
	"strings"
    "sync"
)

var replacer = strings.NewReplacer(".", "", ",", "", "!", "", "?", "", "||", "", "(", "", ")", "", "\"", "", "'", "")
var filewg sync.WaitGroup
var fileChannel = make(chan string, 100)

var Dict = make(map[string]int) // maps tokens to an index
var Dictfile = "dict.csv"
var Matrixfile = "matrix.csv"

/**
  Takes as input a filename containing the full text for a patent on each line. Returns
  a string slice where each entry is the full text of a patent (order is maintained)
  which has had the following transformations applied to it:
  * unescaped HTML sequences
  * lowercase
  * removed all .,!?||()"'
*/
func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
    filewg.Add(1)
    go func() {
        for {
            line, err := reader.ReadString('\n')
            if err != nil {
                //fmt.Println(err)
                fmt.Println("donedoneasdfasdfasdfasdfasdfasf")
                filewg.Done()
                close(fileChannel)
                break
            }
            line = html.UnescapeString(line)
            line = strings.ToLower(line)
            line = strings.Trim(line, " ")
            line = replacer.Replace(line)
            fileChannel <- line
        }
    }()
    filewg.Wait()
	fmt.Println("Finished reading input file", filename)
	//return results
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
func CreateDict(filename string) {
	fmt.Println("Creating token dictionary")
	index := 0
    go readFile(filename)
	for line := range fileChannel {
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
