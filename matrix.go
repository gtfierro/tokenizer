package tokenizer

import (
    "strconv"
    "sync"
    "fmt"
)

var wg sync.WaitGroup

func Generate_chan(lines []string) chan string{
    in := make(chan string)
    go func() {
        defer close(in)
        for _, line := range lines {
            in <- line
        }
    }()
    return in
}


func PCreateMatrix(in <-chan string) {
    patent_index := 0
    for line := range in {
        patent_index += 1
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
    for token, count := range(tmpmap) {
        entry := "(" + strconv.Itoa(patent_index) + "," + strconv.Itoa(Dict[token]) + "," + strconv.Itoa(count) + ")"
        fmt.Println(entry)
    }
}
