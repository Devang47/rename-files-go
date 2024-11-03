package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a directory name")
		return
	}

	dirName := os.Args[1]
	fmt.Println("Reading files from", dirName)

	files, err := os.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			res, err := matchAndConvert(file.Name())
			if err != nil {
				continue
			}

			fmt.Printf("Renaming \"%s\" to \"%s\" \n", file.Name(), res)

			err = os.Rename(fmt.Sprintf("./sample/%s", file.Name()), fmt.Sprintf("./sample/%s", res))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// Takes "file copy 1.txt"
// Returns "file_1.txt"
// or error if no match
func matchAndConvert(filename string) (string, error) {
	pieces := strings.Split(filename, ".")

	// Get the extension
	extn := pieces[len(pieces)-1]

	tmp := strings.Split(pieces[0], " ")
	if (len(tmp) < 2) || (tmp[len(tmp)-2] != "copy") {
		return "", fmt.Errorf("%s no match", filename)
	}

	// extract the name and change positions
	fName := strings.Join(tmp[:len(tmp)-2], "")
	fNum, err := strconv.Atoi(tmp[len(tmp)-1])
	if err != nil {
		return "", fmt.Errorf("%s no match", filename)
	}

	return fmt.Sprintf("%s_%d.%s", fName, fNum, extn), nil
}
