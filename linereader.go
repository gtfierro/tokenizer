package tokenizer

import (
    "bufio"
    "fmt"
    "os"
    "html"
)

func Read_file(filename string) []string{
    results := []string{}
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        results = append(results, html.UnescapeString(line))
    }
    return results
}
