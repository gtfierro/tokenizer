package tokenizer

var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 
                     'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

//var channels = map[byte](chan string){
//    'a' : make(chan string),
//    'b' : make(chan string),
//    'c' : make(chan string),
//    'd' : make(chan string),
//    'e' : make(chan string),
//    'f' : make(chan string),
//    'g' : make(chan string),
//    'h' : make(chan string),
//    'i' : make(chan string),
//    'j' : make(chan string),
//    'k' : make(chan string),
//    'l' : make(chan string),
//    'm' : make(chan string),
//    'n' : make(chan string),
//    'o' : make(chan string),
//    'p' : make(chan string),
//    'q' : make(chan string),
//    'r' : make(chan string),
//    's' : make(chan string),
//    't' : make(chan string),
//    'u' : make(chan string),
//    'v' : make(chan string),
//    'w' : make(chan string),
//    'x' : make(chan string),
//    'y' : make(chan string),
//    'z' : make(chan string),
//    '.' : make(chan string), //catchall channel
//}

var a = make(chan string)
var b = make(chan string)
var c = make(chan string)
var d = make(chan string)
var e = make(chan string)
var f = make(chan string)
var g = make(chan string)
var h = make(chan string)
var i = make(chan string)
var j = make(chan string)
var k = make(chan string)
var l = make(chan string)
var m = make(chan string)
var n = make(chan string)
var o = make(chan string)
var p = make(chan string)
var q = make(chan string)
var r = make(chan string)
var s = make(chan string)
var t = make(chan string)
var u = make(chan string)
var v = make(chan string)
var w = make(chan string)
var x = make(chan string)
var y = make(chan string)
var z = make(chan string)
var all = make(chan string) //catchall channel

var dicts = map[byte](map[string]int){
    'a' : make(map[string]int),
    'b' : make(map[string]int),
    'c' : make(map[string]int),
    'd' : make(map[string]int),
    'e' : make(map[string]int),
    'f' : make(map[string]int),
    'g' : make(map[string]int),
    'h' : make(map[string]int),
    'i' : make(map[string]int),
    'j' : make(map[string]int),
    'k' : make(map[string]int),
    'l' : make(map[string]int),
    'm' : make(map[string]int),
    'n' : make(map[string]int),
    'o' : make(map[string]int),
    'p' : make(map[string]int),
    'q' : make(map[string]int),
    'r' : make(map[string]int),
    's' : make(map[string]int),
    't' : make(map[string]int),
    'u' : make(map[string]int),
    'v' : make(map[string]int),
    'w' : make(map[string]int),
    'x' : make(map[string]int),
    'y' : make(map[string]int),
    'z' : make(map[string]int),
    '.' : make(map[string]int),
}
