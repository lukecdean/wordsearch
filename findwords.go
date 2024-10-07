package main

import (
    "fmt"
    "wordfinder/wordfinder"
) // import

func main() {
    runCriterion("acemnorsuvwxz", 3)
    runCriterion("acemnorsuvwxz", 6)
    runCriterion("acemnorsuvwxz", 8)
    runCriterion("acemnorsuvwxz", 9)
    runCriterion("acemnorsuvwxz", 12)
    runCriterion("acemnorsuvwxz", 16)
} // main()

func runCriterion(charset string, length int) {
    name := fmt.Sprintf("%s_%d", charset, length)
    var criteria []wordfinder.Criterion
    criteria = append(criteria, createLenCriterion(length))
    criteria = append(criteria, createCharSetCriterion(charset))
    wordfinder.Checkwords("wordbank/wordbank.txt", name, criteria)
} // runCriterion()

func createLenCriterion(length int) wordfinder.Criterion {
    return func(word string) bool {
        return islen(length, word)
    }
} // createLenCriterion()

func createCharSetCriterion(charset string) wordfinder.Criterion {
    return func(word string) bool {
        return charsonlyfromset(charset, word)
    }
} // createCharSetCriterion()

func charsonlyfromset(charset string, word string) bool {
    byteset := []byte(charset)
    wordchars := []byte(word)
    for _, char := range wordchars {
        if !sliceContains(byteset, char) {
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

func islen(length int, word string) bool {
    return len(word) == length
} // islen
