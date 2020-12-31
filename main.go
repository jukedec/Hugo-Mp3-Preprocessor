package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"

	id3 "github.com/mikkyang/id3-go"

	"github.com/dhowden/tag"

	// utf32 "golang.org/x/text/encoding/unicode/utf32"

	"io/ioutil"
)

// import "github.com/dhowden/tag"

// import "fmt"

var artistName []string
var n = "\n"
var githubPage = ""

func main() {

	argsWithoutProg := os.Args[1:]

	// arg := os.Args[3]

	// siteName := os.Args[3]

	fmt.Printf("hello, world\n")

	fmt.Println(argsWithoutProg)

	baseDir, err := os.Executable()
	baseDir = baseDir + "/"
	// pushDir, err := os.Executable()

	//TODO: figure out how to send files with second param
	// siteDir := ""
	// if len(argsWithoutProg) > 2 {
	// dir := argsWithoutProg[1]
	// siteDir = argsWithoutProg[1]
	// }

	// pushDir := argsWithoutProg[0]

	fmt.Println("Push dir and Origin Dir:")
	// fmt.Println(baseDir + argsWithoutProg[0])
	// fmt.Println(baseDir + argsWithoutProg[1])

	// pushDir := filepath.Dir(argsWithoutProg[0])
	// dir := filepath.Dir(argsWithoutProg[1])

	dir := filepath.Clean(argsWithoutProg[0])
	pushDir := filepath.Clean(argsWithoutProg[1])
	baseURL := argsWithoutProg[2]

	fmt.Println("baseURL:")
	fmt.Println(baseURL)


	// pushDir := baseDir + argsWithoutProg[0]
	// dir := baseDir + argsWithoutProg[1]

	fmt.Println("FILEPATH SET AS: Origin dir and Push Dir:")
	fmt.Println(dir)
	fmt.Println(pushDir)
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
	folderPath := filepath.Join(pushDir, "/static/")

	fmt.Println("Making directory: " + folderPath)
	os.MkdirAll(folderPath, os.ModePerm)
	artistName := ""
	for _, file := range files {
		fmt.Println(file)
		fmt.Println(reflect.TypeOf(file).String())
		last3 := file[len(file)-4:]
		if strings.ToLower(last3) != ".mp3" {
			// fmt.Printf("\n")
			fmt.Println("GETTING SOMETHING WEIRD NOT MP3 for: " + file)

		} else {

			origFileSlice := strings.Split(file, "/")
			origFileName := origFileSlice[len(origFileSlice)-1]
			var staticDir = filepath.Join(folderPath, origFileName)
			artistName = makeMd(file, pushDir, folderPath)
			fmt.Println("SETTING TO ARTIST: " + artistName + "/n")

			copy(file, staticDir)

			makeConfig(pushDir, baseURL)
		}
	}

	// fmt.Printf(mp3File.Title())

	// 	---
	// title: "My First Post toast"
	// date: 2019-04-21T14:40:44-05:00
	// draft: false
	// ---

	// Firsty
	fmt.Println("theArtistName=")
	fmt.Println(getBaseSiteName(artistName))

}

func check(e error) {

	if e != nil {
		fmt.Println("CHECKED AND GOT ERROR")

		panic(e)
	}
}

func getBaseSiteName(artistName string) string {
	fileReg, err := regexp.Compile("[^a-zA-Z0-9]+")
	check(err)
	name := fileReg.ReplaceAllString(artistName, "")
	reg, err := regexp.Compile("[^ -~]+")
	check(err)
	nameStripped := reg.ReplaceAllString(name, "")
	return nameStripped
}

func makeMd(f string, dir string, staticDir string) string {
	fmt.Printf("MAKING MD from:" + f + "\n")
	mp3File, err := id3.Open(f)

	// These methods are for Title, Artist, Album, Year, Genre, and Comments.
	lyricsFrame := mp3File.Frame("USLT")

	lyrics := ""

	if lyricsFrame != nil {
		lyrics = lyricsFrame.String()
	} else {
	}

	var mp3Title = mp3File.Title()
	artistName = append(artistName, mp3File.Artist())
	fmt.Printf("mp3Title: " + mp3Title + "\n")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n")

	fileReg, err := regexp.Compile("[^a-zA-Z0-9]+")
	reg, err := regexp.Compile("[^ -~]+")

	check(err)
	mp3TitleStripped := reg.ReplaceAllString(mp3Title, "")
	mp3YearStripped := reg.ReplaceAllString(mp3File.Year(), "")

	fmt.Printf("mp3TitleStripped: " + mp3TitleStripped + "\n")

	const (
		layoutISO = "2006"
		layoutUS  = "January 2, 2006"
	)

	fmt.Printf("staticDir: " + staticDir + "\n")
	fmt.Printf("dir: " + dir + "\n")

	mdFile := fileReg.ReplaceAllString(mp3Title, "") + ".md"
	mdPath := filepath.Join(dir, "content/posts")
	os.MkdirAll(mdPath, os.ModePerm)
	fmt.Printf("FILE IS: " + mdFile + "\n")

	mdPath = filepath.Join(mdPath, mdFile)
	origFileSlice := strings.Split(f, "/")
	origFileName := origFileSlice[len(origFileSlice)-1]
	// origFileName = "/" + origFileName
	fmt.Printf("origFileName: " + origFileName + "\n")

	mp3Source := "/" + getBaseSiteName(mp3File.Artist()) + "/" + origFileName
	// mp3Source := "/" + origFileName

	fmt.Printf("mp3Source: " + mp3Source + "\n")

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

	output.WriteString("\"" + mp3YearStripped + "-01-01\"")
	output.WriteString(n)

	output.WriteString("---")
	output.WriteString(n)
	output.WriteString("<audio controls>")

	output.WriteString("<source src='")

	// mp3Source := filepath.Join(githubPage, origFileName)

	output.WriteString(mp3Source)
	output.WriteString("'  type='audio/mpeg'>")
	output.WriteString(n)
	output.WriteString("</audio>")
	output.WriteString(n)
	output.WriteString(lyrics)
	output.WriteString(n)

	fmt.Println(n + "md files are being saved to: " + mdPath)
	writeErr := ioutil.WriteFile(mdPath, output.Bytes(), 0777)
	check(writeErr)

	getImg(f, staticDir)

	// reflect.TypeOf()

	return mp3File.Artist()
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

func makeConfig(dir string, baseURL string) {
	var output bytes.Buffer

	//Remove because I haven't figured out solid deploy yet.
	fileReg, err := regexp.Compile("[^a-zA-Z0-9]+")
	check(err)
	name := fileReg.ReplaceAllString(artistName[0], "")
	reg, err := regexp.Compile("[^ -~]+")
	check(err)
	nameStripped := reg.ReplaceAllString(name, "")
	// nameStripped := "musicMonth"
	fmt.Println(nameStripped)

	fmt.Println("WRITING THE CONFIG:")

	// githubPage = "http://frigginglorious.github.io/" + nameStripped
	// localPage := "http://localhost:1313/" + nameStripped
	// localPage := "http://play.jukedec.com/" + nameStripped

	localPage := baseURL + "/" + nameStripped



	// output.WriteString("baseURL = \"" + githubPage + "/\"" + n)
	// output.WriteString("baseURL = \"" + "http://localhost:1313" + "/\"" + n)
	// output.WriteString("baseURL = \"" + "/\"" + n)

	output.WriteString("baseURL = \"" + localPage + "/\"" + n)

	output.WriteString("languageCode = \"en-us\"" + n)
	output.WriteString("title = \"" + nameStripped + "\"" + n)

	// Just import all themes as the name "import"
	output.WriteString("theme = \"import\"" + n)

	output.WriteString("style = \"default\"" + n)
	output.WriteString("[params]" + n + "authorimage = \"cover.jpg\"" + n)
	output.WriteString("dateformat = \"2006\"" + n)

	// [markup.goldmark.renderer]
	// unsafe = true

	output.WriteString("[markup.goldmark.renderer]" + n)
	output.WriteString("unsafe = true" + n)

	fmt.Println("WRITING to URL:" + localPage)

	configFile := filepath.Join(dir, "config.toml")
	fmt.Println("WRITING THE CONFIG TO:" + configFile)

	writeErr := ioutil.WriteFile(configFile, output.Bytes(), 0777)
	check(writeErr)
}

func getImg(f string, staticDir string) {

	readFile, err := os.Open(f)
	check(err)
	m, err := tag.ReadFrom(readFile)
	check(err)
	fmt.Println("ID3 Tag Title?:")
	fmt.Println(m.Title() + "\n")

	picture := m.Picture()
	fmt.Println("Type of picture?:")

	fmt.Println(reflect.TypeOf(picture))
	fmt.Println("Val of picture?:")
	fmt.Println(picture)

	if picture != nil {
		fmt.Println("SAVING PIC:")
		img, _, _ := image.Decode(bytes.NewReader(picture.Data))

		var opt jpeg.Options

		opt.Quality = 80
		// ok, write out the data into the new JPEG file

		out, err := os.Create(staticDir + "/cover.jpg")
		check(err)
		err = jpeg.Encode(out, img, &opt)
	} else {
		fmt.Println("NO PIC:")

	}

	log.Print(m.Format()) // The detected format.
	log.Print(m.Title())  // The title of the track (see Metadata interface for more details).

}
