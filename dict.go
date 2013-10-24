package tokenizer

import (
	"fmt"
	_ "html"
	_ "strings"
    "bytes"
	"sync"
    _ "runtime"
)

var tokenChannel = make(chan []string)
var doneChannel = make(chan bool)
var tokenwg sync.WaitGroup

func process() {
	index := 0
	for {
		select {
		case tokens := <-tokenChannel:
			for _, token := range tokens {
                token := string(token)
				if Dict[token] == 0 {
					Dict[token] = index
					index += 1
				}
			}
		case <-doneChannel:
			break
		}
	}
}

func deliver(line []byte) {
	defer tokenwg.Done()
	//line = html.UnescapeString(line)
	line = bytes.ToLower(line)
	line = bytes.Trim(line, " ")
    line = bytes.Replace(line, []byte("."), []byte(""), -1)
    line = bytes.Replace(line, []byte(","), []byte(""), -1)
    line = bytes.Replace(line, []byte("!"), []byte(""), -1)
    line = bytes.Replace(line, []byte("?"), []byte(""), -1)
    line = bytes.Replace(line, []byte("|"), []byte(""), -1)
    line = bytes.Replace(line, []byte("("), []byte(""), -1)
    line = bytes.Replace(line, []byte(")"), []byte(""), -1)
    line = bytes.Replace(line, []byte("'"), []byte(""), -1)
    line = bytes.Replace(line, []byte("\""),[]byte(""), -1)
    line = bytes.Replace(line, []byte("\t"),[]byte(""), -1)
    line = bytes.Replace(line, []byte("\n"),[]byte(""), -1)
	tokens := tokenize(line, false)
	tokenChannel <- tokens
}

/**
  Given the output of Read_file, populates Dict, a map[string]int.
  Iterates through each of the found lines, tokenizes the line,
  and adds tokens to the Dict, which maintains a mapping of a token
  to its index
*/
func CreateDict(filename string) {
	go process()
	go readFile(filename)
	fmt.Println("Creating token dictionary")
	for line := range fileChannel {
		tokenwg.Add(1)
		go deliver(line)
	}
	tokenwg.Wait()
	doneChannel <- true
	fmt.Println("Finished creating token dictionary with", len(Dict), "items")
}
