package tokenizer

import (
    "sync"
    "strings"
	"html"
    "fmt"
)


var tokenChannel = make(chan []string)
var doneChannel = make(chan bool)
var tokenwg sync.WaitGroup

func process() {
    index := 0
    for {
        select {
        case tokens := <- tokenChannel:
            for _, token := range tokens {
                if Dict[token] == 0 {
                    Dict[token] = index
                    index += 1
                }
            }
        case <- doneChannel:
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

func CreateDict2(lines []string) {
    go process()
	fmt.Println("Creating token dictionary")
    for _, line := range lines {
        tokenwg.Add(1)
        go deliver(line)
    }
    tokenwg.Wait()
    fmt.Println("here")
    doneChannel <- true
	fmt.Println("Finished creating token dictionary with", len(Dict), "items")
}
