package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	id3 "github.com/mikkyang/id3-go"

	// utf32 "golang.org/x/text/encoding/unicode/utf32"

	"io/ioutil"
)

// import "fmt"

func main() {
	fmt.Printf("hello, world\n")

	// fixUtf := func(r rune) rune {
	// 	if r == utf8.RuneError { // && r != ' '
	// 		// fmt.Printf(r)
	// 		fmt.Println(r, string(r), strconv.QuoteRune(r))
	// 		fmt.Printf("\n")

	// 		return -1
	// 	}
	// 	return r
	// }

	var files []string

	// root, err := filepath.Abs(filepath.Dir(os.Args[0])) // filepath.Dir(os.Args[0])
	dir, err := os.Getwd()
	check(err)
	fmt.Printf(dir)

	pathErr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	check(pathErr)
	for _, file := range files {
		fmt.Println(file)
		fmt.Println(reflect.TypeOf(file).String())

		makeMd(file)
	}

	// fmt.Printf(mp3File.Title())

	// 	---
	// title: "My First Post toast"
	// date: 2019-04-21T14:40:44-05:00
	// draft: false
	// ---

	// Firsty

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeMd(f string) bool {

	last3 := f[len(f)-4:]
	if string(last3) != ".mp3" {
		return false
	}

	mp3File, err := id3.Open(f)
	var mp3Title = mp3File.Title()
	fmt.Printf(mp3Title)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n")

	fileReg, err := regexp.Compile("[^a-zA-Z0-9]+")
	reg, err := regexp.Compile("[^ -~]+")

	check(err)
	mp3TitleStripped := reg.ReplaceAllString(mp3Title, "")
	mp3YearStripped := reg.ReplaceAllString(mp3File.Year(), "")

	const (
		layoutISO = "2006"
		layoutUS  = "January 2, 2006"
	)

	mdFile := fileReg.ReplaceAllString(mp3Title, "") + ".md"

	origFileSlice := strings.Split(f, "/")
	origFileName := origFileSlice[len(origFileSlice)-1]
	fmt.Printf(origFileName)

	// date := mp3YearStripped
	// t, _ := time.Parse(layoutISO, date)

	// t1, e := time.Parse(
	// time.RFC3339,
	// mp3YearStripped)

	// check(e)

	// f, err := os.Create("gospel.md")
	// check(err)
	// defer f.Close()
	var output bytes.Buffer
	var n = "\n"
	output.WriteString("---")
	output.WriteString(n)
	output.WriteString("title: ")
	output.WriteString("\"")
	// output.WriteString(strings.Map(fixUtf, mp3Title))
	output.WriteString(mp3TitleStripped)

	output.WriteString("\"")
	output.WriteString(n)
	output.WriteString("draft: false")

	output.WriteString(n)
	output.WriteString("date: ")

	output.WriteString(mp3YearStripped)
	output.WriteString(n)

	output.WriteString("---")
	output.WriteString(n)
	output.WriteString("<audio controls>")

	output.WriteString("<source src='")
	output.WriteString(origFileName)
	output.WriteString("'  type='audio/mpeg'>")
	output.WriteString(n)
	output.WriteString("</audio>")
	output.WriteString(n)

	writeErr := ioutil.WriteFile(mdFile, output.Bytes(), 0644)
	check(writeErr)
	return true
}
