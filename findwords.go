package main

import "wordfinder/wordfinder"

func main() {
    wordfinder.Checkwords("wordbank/wordbank.txt", "output/ishexand7chars.txt", ishexand7chars)
} // main()

// Wordfitscriteria functions ie funcs that define a certain criteria for words
func ishex(word string) bool {
    hexchars := []byte("abcdef")
    return charsonlyfromset(hexchars, word)
} // ishex()

func ishexand6chars(word string) bool {
    hexchars := []byte("abcdef")
    return charsonlyfromset(hexchars, word) && islen(word, 6)
} // ishexand8chars()

func ishexand7chars(word string) bool {
    hexchars := []byte("abcdef")
    return charsonlyfromset(hexchars, word) && islen(word, 7)
} // ishexand8chars()

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

func sliceContains(bytes []byte, b byte) bool {
    for _, slicebyte := range bytes {
        if slicebyte == b {
            return true
        } // if
    } // for b
    return false
} // sliceContains()

func islen(word string, length int) bool {
    return len(word) == length
} // islen
