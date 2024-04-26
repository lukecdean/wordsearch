package main

import (
    "bufio"
    "fmt"
    "os"
) // import

func main() {
    filepath := "words.txt"
    //var wordcount int = 10000
    file, err := os.Open(filepath)
    if err != nil {
        fmt.Println("err opening file:", err)
    } // if err
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var words []string
    //for i := 0; i < wordcount; i++ {
        //scanner.Scan()
        //words = append(words, scanner.Text())
    //} // for scanner
    for scanner.Scan() {
        words = append(words, scanner.Text())
    } // for 

    all(words)
} // main()

func all(words []string) {
    for _, word := range words {
        if canBeRepresentedInHex(word) {
            fmt.Println(word)
        } // if
    } // for
} // all()

func canBeRepresentedInHex(word string) bool {
    hexchars := []byte{'a', 'b', 'c', 'd', 'e', 'o', 'l', 's'}
    wordchars := []byte(word)
    for _, char := range wordchars {
        if !sliceContains(hexchars, char) {
            return false
        } // if
    } // for char
    return true
} // canBeHex()

func sliceContains(bytes []byte, b byte) bool {
    for _, slicebyte := range bytes {
        if slicebyte == b {
            return true
        } // if
    } // for b
    return false
} // sliceContains()
