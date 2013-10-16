package tokenizer

import (
    "bufio"
    "fmt"
    "os"
    "html"
    "strings"
	"github.com/agonopol/go-stem/stemmer"
)

var replacer = strings.NewReplacer(".", "", ",", "", "!", "", "?", "", "||", "", "(", "", ")", "")

var Dict = map[string]int // maps tokens to an index
var Dictfile = "dict.csv"
var Matrixfile = "matrix.csv"

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

/**
    Takes as input a space-delimited string and returns a slice of all the
    tokens in that string.

    If the `stem` flag is True, applies the Porter stemming algorithm
    to each token
*/
func tokenize(instring string, stem bool) []string {
    tokens := strings.Split(instring, " ")
    if stem {
        res := []string{}
        for _, t := range tokens {
            stem := stemmer.Stem([]byte(t))
            res = append(res, string(stem))
        }
        return res
    }
    return tokens
}
