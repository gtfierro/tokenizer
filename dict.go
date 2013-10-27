package tokenizer

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var tokenChannel = make(chan [][20]byte)
var doneChannel = make(chan bool)
var tokenwg sync.WaitGroup

type Entry struct {
	token [20]byte
	index int
}

var entryChannel = make(chan *Entry)

func process() {
	index := 0
	for {
		select {
		case tokens := <-tokenChannel:
			for _, token := range tokens {
				if Dict[token] == 0 {
					Dict[token] = index
					e := &Entry{token, index}
					entryChannel <- e
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
	line = UnescapeBytes(line)
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
	line = bytes.Replace(line, []byte("”"), []byte(""), -1)
	line = bytes.Replace(line, []byte("“"), []byte(""), -1)
	line = bytes.Replace(line, []byte("\""), []byte(""), -1)
	line = bytes.Replace(line, []byte("\t"), []byte(""), -1)
	line = bytes.Replace(line, []byte("\n"), []byte(""), -1)
	tokens := tokenize(line, false)
	tokenChannel <- tokens
}

func outputDict() {
	outfile, err := os.Create(Dictfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()
	writer := bufio.NewWriter(outfile)
	for e := range entryChannel {
		writer.WriteString(string(e.token[:]) + "," + strconv.Itoa(e.index) + "\n")
	}
	writer.Flush()
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
	go outputDict()
	fmt.Println("Creating token dictionary")
	for line := range fileChannel {
		tokenwg.Add(1)
		go deliver(line)
	}
	tokenwg.Wait()
	doneChannel <- true
	fmt.Println("Finished creating token dictionary with", len(Dict), "items")
}
