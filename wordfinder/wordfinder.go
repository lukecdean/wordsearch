package wordfinder

import (
    "bufio"
    "fmt"
    "os"
) // import

type Wordfitscriteria func(string) bool

func Checkwords(wordbankpath string, outputfilename string, fitscriteria Wordfitscriteria) {
    // create output file
    outputfile := openoutputfile(outputfilename)
    defer outputfile.Close()

    // load the words from the wordbank into a list
    words := getwords(wordbankpath)

    // search the words
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

func openoutputfile(outputfilename string) *os.File {
    outputdir := "output/"
    outputfile, err := os.Create(outputdir + outputfilename + ".txt")
    if err != nil {
        fmt.Println("err creating file:", err)
    } // if err
    return outputfile
} // openoutfile()
