package tokenizer

import (
	"fmt"
	"strconv"
	"sync"
)

/** Parallel **/

var wg sync.WaitGroup

func PCreateMatrix(lines []string) {
	for patent_index, line := range lines {
		wg.Add(1)
		go pemit_sparse(patent_index, line)
	}
	wg.Wait()
}

func pemit_sparse(patent_index int, line string) {
	defer wg.Done()
	tmpmap := make(map[string]int)
	for _, token := range tokenize(line, false) {
		tmpmap[token] += 1
	}
	for token, count := range tmpmap {
		entry := "(" + strconv.Itoa(patent_index) + "," + strconv.Itoa(Dict[token]) + "," + strconv.Itoa(count) + ")"
		fmt.Println(entry)
	}
}

/** Sequential **/

func CreateMatrix(lines []string) {
	for patent_index, line := range lines {
		emit_sparse(patent_index, line)
	}
}

func emit_sparse(patent_index int, line string) {
	tmpmap := make(map[string]int)
	for _, token := range tokenize(line, false) {
		tmpmap[token] += 1
	}
	for token, count := range tmpmap {
		entry := "(" + strconv.Itoa(patent_index) + "," + strconv.Itoa(Dict[token]) + "," + strconv.Itoa(count) + ")"
		fmt.Println(entry)
	}
}
