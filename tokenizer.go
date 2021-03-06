package tokenizer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/agonopol/go-stem/stemmer"
	"os"
	"sync"
)

// synchronizes the goroutine that reads the input file
var filewg sync.WaitGroup

var Dict = make(map[[20]byte]int)

// output files
var Dictfile = "dict.csv"
var Matrixfile = "matrix.csv"

/**
  Takes as input a filename containing the full text for a patent on each line and the
  deliver channel over which the read lines will be delivered. Delivers each line as
  a byte slice.
*/
func readFile(filename string, channel chan []byte) {
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
				close(channel)
				break
			}
			channel <- line
		}
	}()
	filewg.Wait()
	fmt.Println("Finished reading input file", filename)
}

/*
Truncates a byte slice to length 20 so that it can be
used as a key in the Dict map
*/
func slice2array(in []byte) [20]byte {
	out := [20]byte{}
	for i, char := range in {
		if i == 20 {
			break
		}
		out[i] = char
	}
	return out
}

/*
   Returns true if the input byte slice
   is all digits
*/
func isDigits(target []byte) bool {
	all := true
	for _, b := range target {
		if b <= '0' || '9' <= b {
			all = false
		}
	}
	return all
}

/**
  Takes as input a space-delimited string and returns a slice of all the
  tokens in that string.

  If the `stem` flag is True, applies the Porter stemming algorithm
  to each token
*/
func tokenize(instring []byte, stem bool) [][20]byte {
	tokens := bytes.Split(instring, []byte(" "))
	res := [][20]byte{}
	for _, t := range tokens {
		if isDigits(t) {
			continue
		}
		if stem {
			t = stemmer.Stem(t)
		}
		res = append(res, [20]byte(slice2array(t)))
	}
	return res
}
