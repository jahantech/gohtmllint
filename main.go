package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("No Arguments specified")
		fmt.Println("Usage (Files): gohtmllint test.html")
		fmt.Println("Usage (Folder): gohtmllint TestDir")
		return
	}
	Check_File_Folder := os.Args[1]

	HTMLOnly := flag.Bool("HTMLOnly", true, "Set this to false if you want to scan all the files")
	BasicHtmlTagChecker(Check_File_Folder, *HTMLOnly)
}

func BasicHtmlTagChecker(Check_File_Folder string, HTMLOnly bool) {
	if len(Check_File_Folder) < 1 {
		fmt.Println("No File/Folder name was specified")
		return
	}
	fdir, err := os.Open(Check_File_Folder)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer fdir.Close()

	Tag, err := regexp.Compile("<.*?>")

	CloseTag, _ := regexp.Compile(`<\/.*?>`)
	fileinfo, err := fdir.Stat()

	if err != nil {
		fmt.Println("Unable to get File/Folder Stat")
		fmt.Println(err)
		return
	}

	switch mode := fileinfo.Mode(); {
	case mode.IsDir():
		return
	case mode.IsRegular():
		content, err := ioutil.ReadFile(Check_File_Folder)

		if err != nil {
			fmt.Println("Unable to read the file")
			return
		}
		opentags := 0
		closedtags := 0
		AllTags := Tag.FindAllString(string(content), -1)
		for _, v := range AllTags {

			if CloseTag.Match([]byte(v)) {
				fmt.Println("Closing Tag: " + v)
				opentags = opentags + 1

			} else {
				fmt.Println("Opening Tag: " + v)
				closedtags = closedtags + 1
			}
		}

		if closedtags == opentags {
			fmt.Println("Passed: " + Check_File_Folder)
		} else {
			fmt.Println("Failed - Count mismatch: " + Check_File_Folder)
		}
		return
	}
}
