package main

import (
    "bufio"
    "fmt"
    "os"
) // import

type wordfitscriteria func(string) bool

func main() {
    words := getwords("words.txt")
    checkwords(words, "output/ishexexpanded.txt", ishexexpanded)
} // main()

func getwords(wordbankpath string) []string {
    // open the file
    wordbankfile := openwordbank(wordbankpath)
    defer wordbankfile.Close()

    // scan through the file to load into a list
    scanner := bufio.NewScanner(wordbankfile)
    var words []string
    for scanner.Scan() {
        words = append(words, scanner.Text())
    } // for 
    return words
} // getwords()

func openwordbank(wordbankpath string) *os.File {
    wordbank, err := os.Open(wordbankpath)
    if err != nil {
        fmt.Println("err opening file:", err)
    } // if err
    return wordbank
} // openwordbank()

func openoutputfile(outputfilepath string) *os.File {
    outputfile, err := os.Create(outputfilepath)
    if err != nil {
        fmt.Println("err creating file:", err)
    } // if err
    return outputfile
} // openoutfile()

func checkwords(words []string, outputfilepath string, fitscriteria wordfitscriteria) {
    outputfile := openoutputfile(outputfilepath)
    defer outputfile.Close()

    for _, word := range words {
        if fitscriteria(word) {
            //fmt.Println(word)
            _, err := outputfile.WriteString(word + "\n")
            if err != nil {
                fmt.Println("err writing to file:", err)
            } // err
        } // if
    } // for
} // all()

func ishex(word string) bool {
    hexchars := []byte("abcdef")
    return charsonlyfromset(hexchars, word)
} // ishex()

func ishexand8chars(word string) bool {
    hexchars := []byte("abcdef")
    return charsonlyfromset(hexchars, word) && islen(word, 8)
} // ishexand8chars()

func ishexexpanded(word string) bool {
    hexchars := []byte("abcdefslo")
    return charsonlyfromset(hexchars, word)
} // ishexexpanded()

func ishexexpandedand8chars(word string) bool {
    hexchars := []byte("abcdefslo")
    return charsonlyfromset(hexchars, word) && islen(word, 8)
} // ishexexpanded()

func charsonlyfromset(charset []byte, word string) bool {
    wordchars := []byte(word)
    for _, char := range wordchars {
        if !sliceContains(charset, char) {
            return false
        } // if
    } // for char
    return true
} // canBeHex()

func islen(word string, length int) bool {
    return len(word) == length
} // islen

func sliceContains(bytes []byte, b byte) bool {
    for _, slicebyte := range bytes {
        if slicebyte == b {
            return true
        } // if
    } // for b
    return false
} // sliceContains()
