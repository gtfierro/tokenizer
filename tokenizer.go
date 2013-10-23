package tokenizer

import (
	"bufio"
	"fmt"
    "bytes"
	"github.com/agonopol/go-stem/stemmer"
	_ "html"
	"os"
	_ "strings"
	"sync"
)

var filewg sync.WaitGroup
var fileChannel = make(chan []byte)

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
			line, err := reader.ReadBytes('\n')
			if err != nil {
				filewg.Done()
				close(fileChannel)
				break
			}
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
func tokenize(instring []byte, stem bool) [][]byte {
	tokens := bytes.Split(instring, []byte(" "))
	if stem {
		res := [][]byte{}
		for _, t := range tokens {
			stem := stemmer.Stem(t)
			res = append(res, stem)
		}
		return res
	}
	return tokens
}
