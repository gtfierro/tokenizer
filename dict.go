package tokenizer

import (
    "fmt"
	"sync"
    "strings"
)

var pwg sync.WaitGroup

func dispatch_channels() {
    go make_dict('a', a)
    go make_dict('b', b)
    go make_dict('c', c)
    go make_dict('d', d)
    go make_dict('e', e)
    go make_dict('f', f)
    go make_dict('g', g)
    go make_dict('h', h)
    go make_dict('i', i)
    go make_dict('j', j)
    go make_dict('k', k)
    go make_dict('l', l)
    go make_dict('m', m)
    go make_dict('n', n)
    go make_dict('o', o)
    go make_dict('p', p)
    go make_dict('q', q)
    go make_dict('r', r)
    go make_dict('s', s)
    go make_dict('t', t)
    go make_dict('u', u)
    go make_dict('v', v)
    go make_dict('w', w)
    go make_dict('x', x)
    go make_dict('y', y)
    go make_dict('z', z)
    go make_dict('.', all)
}


func make_dict(letter byte, channel chan string) (byte, int, map[string]int) {
    //dict := dicts[letter]
    i := 0
    for {
        token := <- channel
        if dicts[letter][token] == 0 {
            dicts[letter][token] = i
            i += 1
        }
    }

    return letter, len(dicts[letter]), dicts[letter]
}

func dispatch(line string) {
    defer pwg.Done()
    tokens := tokenize(line, false)
    for _, token := range tokens {
        // check if token is only whitespace
        if len(strings.Replace(token, " ", "", -1)) > 0 {
            first := []byte(token)[0] // get first char of token
            switch {
            case first == 'a':
                a <- token
            case first == 'b':
                b <- token
            case first == 'c':
                c <- token
            case first == 'd':
                d <- token
            case first == 'e':
                e <- token
            case first == 'f':
                f <- token
            case first == 'g':
                g <- token
            case first == 'h':
                h <- token
            case first == 'i':
                i <- token
            case first == 'j':
                j <- token
            case first == 'k':
                k <- token
            case first == 'l':
                l <- token
            case first == 'm':
                m <- token
            case first == 'n':
                n <- token
            case first == 'o':
                o <- token
            case first == 'p':
                p <- token
            case first == 'q':
                q <- token
            case first == 'r':
                r <- token
            case first == 's':
                s <- token
            case first == 't':
                t <- token
            case first == 'u':
                u <- token
            case first == 'v':
                v <- token
            case first == 'w':
                w <- token
            case first == 'x':
                x <- token
            case first == 'y':
                y <- token
            case first == 'z':
                z <- token
            default:
                all <- token
            }
        }
    }
}
                   
/**
We want to create a dictionary of map[string]int that maps a token
to a unique index. This will enable the creation of a sparse matrix
representation of the full texts

For each line in the array of all lines, we want to dispatch a goroutine
that calls the `tokenize` method (see tokenizer.go), and then for each word
in the line, sends the word to the correct mapmaker. Each map maker should
run in its own goroutine and should accept words over a buffered input channel.

Args:
lines []string: each `line` in lines is a space-separated list of all tokens
                for a given patent.

*/

func PCreateDict(lines []string) {
    fmt.Println("Creating token dictionary")
    dispatch_channels()
    for _, line := range lines {
        pwg.Add(1)
        go dispatch(line)
    }
    pwg.Wait()
    for letter, dict := range dicts {
        fmt.Println(string(letter), len(dict))
    }
}
