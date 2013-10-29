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
var fileChannel = make(chan []byte)
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

/* loops through `target` and replaces
each instance of `current` with `replacement`
*/
func replace(target []byte, current, replacement byte) []byte {
    for i, b := range target {
        if b == current {
            target[i] = replacement
        }
    }
    return target
}

/* loops through `target` and removes all
instances of `current`
*/

func remove(target []byte, current byte) []byte {
    i1 := 0
    i2 := 0
    length := len(target)
    for {
        if target[i2] != current {
            target[i1] = target[i2]
            i1 += 1
        }
        i2 += 1
        if i2 == length {
            break
        }
    }
    return target[:i1]
}

func deliver(line []byte) {
	line = UnescapeBytes(line)
	line = bytes.ToLower(line)
	line = bytes.Trim(line, " ")
	line = remove(line, '.')
	line = remove(line, ',')
	line = remove(line, '!')
	line = remove(line, '?')
	line = remove(line, '|')
	line = remove(line, '(')
	line = remove(line, ')')
	line = remove(line, '"')
	//line = remove(line, '”')
	//line = remove(line, '“')
	line = remove(line, '\'')
	line = remove(line, '\t')
	line = remove(line, '\n')
	line = replace(line, '/', ' ')
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
		n := bytes.Index(e.token[:], []byte{0})
		if n < 0 {
			n = 20
		}
		writer.WriteString(string(e.token[:n]) + "," + strconv.Itoa(e.index) + "\n")
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
	go readFile(filename, fileChannel)
	go outputDict()
	fmt.Println("Creating token dictionary")
    for i := 0; i < 100; i++ {
        tokenwg.Add(1)
        go func() {
            for line := range fileChannel {
                deliver(line)
            }
            tokenwg.Done()
        }()
    }
	tokenwg.Wait()
	doneChannel <- true
	fmt.Println("Finished creating token dictionary with", len(Dict), "items")
}
