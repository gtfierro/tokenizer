package tokenizer

import (
    "bufio"
    "fmt"
    "os"
    "html"
    "strings"
)

var replacer = strings.NewReplacer(".", "", ",", "", "!", "", "?", "", "||", "", "(", "", ")", "")

/**
    Takes as input a filename containing the full text for a patent on each line. Returns
    a string slice where each entry is the full text of a patent (order is maintained)
    which has had the following transformations applied to it:
    * unescaped HTML sequences
    * lowercase
    * removed all .,!?||()
*/
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
