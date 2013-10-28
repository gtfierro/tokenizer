package tokenizer


import (
	"fmt"
	"strconv"
	"sync"
    "os"
    "bufio"
)

var wg sync.WaitGroup
var matrixFileChannel = make(chan []byte)

var in = make(chan string, 100)

func outputMatrix() {
	outfile, err := os.Create(Matrixfile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outfile.Close()
	writer := bufio.NewWriter(outfile)
	for e := range in {
		writer.WriteString(e)
	}
	writer.Flush()
}


func printMap(patent_index int, tmpMap map[[20]byte]int) {
	for token, count := range tmpMap {
		entry := "(" + strconv.Itoa(patent_index) + "," + strconv.Itoa(Dict[token]) + "," + strconv.Itoa(count) + ")\n"
		in <- entry
	}
}

func CreateMatrix(filename string) {
    go outputMatrix()
    go readFile(filename, matrixFileChannel)
    fmt.Println("Creating sparse matrix")
    patentIndex := 0
    for line := range matrixFileChannel {
        wg.Add(1)
        go emitSparse(patentIndex, line)
        patentIndex += 1
    }
    wg.Wait()
}

func emitSparse(patentIndex int, line []byte) {
    defer wg.Done()
    tmpMap := make(map[[20]byte]int)
    for _, token := range tokenize(line, false) {
        tmpMap[token] += 1
    }
    printMap(patentIndex, tmpMap)
}
