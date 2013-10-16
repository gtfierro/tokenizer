package tokenizer

import (
    "bufio"
    "fmt"
    "os"
    "html"
    "strings"
)

var replacer = strings.NewReplacer(".", "", ",", "", "!", "", "?", "", "||", "", "(", "", ")", "")

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
        line = html.UnescapeString(line)
        line = strings.ToLower(line)
        line = replacer.Replace(line)
        results = append(results, line)
    }
    return results
}
