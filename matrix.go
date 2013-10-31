package tokenizer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
		entry := strconv.Itoa(patent_index) + "," + strconv.Itoa(Dict[token]) + "," + strconv.Itoa(count) + "\n"
		in <- entry
	}
}
