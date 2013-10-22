package tokenizer

import (
	"fmt"
	"html"
	"strings"
	"sync"
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

func deliver(line string) {
	defer tokenwg.Done()
	line = html.UnescapeString(line)
	line = strings.ToLower(line)
	line = strings.Trim(line, " ")
	line = replacer.Replace(line)
	tokens := tokenize(line, false)
	tokenChannel <- tokens
}

func CreateDict2(filename string) {
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
