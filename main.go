package main

import (
	"bytes"
	"fmt"
	"io"
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

var artistName []string
var n = "\n"

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// arg := os.Args[3]

	fmt.Printf("hello, world\n")
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	// fmt.Println(arg)

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
	dir, err := os.Executable()
	dir = filepath.Dir(dir)
	check(err)
	fmt.Printf("GETTING FILES IN DIR:\n")
	fmt.Printf(dir)
	fmt.Printf("\n")

	// os.Exit(1)

	pathErr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	check(pathErr)
	// folderPath := dir + "/band/static/"
	folderPath := filepath.Join(dir, "/static/")

	fmt.Println("Making directory: " + folderPath)
	os.MkdirAll(folderPath, os.ModePerm)
	for _, file := range files {
		fmt.Println(file)
		fmt.Println(reflect.TypeOf(file).String())
		last3 := file[len(file)-4:]
		if strings.ToLower(last3) != ".mp3" {
			// fmt.Printf("\n")
			fmt.Println("GETTING SOMETHING WEIRD NOT MP3")

		} else {
			makeMd(file)
			makeConfig()
			origFileSlice := strings.Split(file, "/")
			origFileName := origFileSlice[len(origFileSlice)-1]
			var staticDir = filepath.Join(folderPath, origFileName)
			copy(file, staticDir)
		}

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

	mp3File, err := id3.Open(f)
	var mp3Title = mp3File.Title()
	artistName = append(artistName, mp3File.Artist())
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
	// origFileName = "/" + origFileName
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

func copy(src, dst string) (int64, error) {
	fmt.Println(n + "Moving file to " + dst)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func makeConfig() {
	var output bytes.Buffer

	fileReg, err := regexp.Compile("[^a-zA-Z0-9]+")
	check(err)
	name := fileReg.ReplaceAllString(artistName[0], "")
	githubPage := "http://frigginglorious.github.io/" + name
	output.WriteString("baseURL = \"" + githubPage + "\"" + n)
	output.WriteString("languageCode = \"en-us\"" + n)
	output.WriteString("title = \"" + artistName[0] + "\"" + n)
	output.WriteString("theme = \"hyde\"" + n)
	output.WriteString("style = \"default\"" + n)

	configFile := "config.toml"
	writeErr := ioutil.WriteFile(configFile, output.Bytes(), 0644)
	check(writeErr)
}
