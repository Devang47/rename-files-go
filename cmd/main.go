package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	// read all files from "sample" directory
	files, err := os.ReadDir("./sample")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			res, err := matchAndConvert(file.Name())
			if err != nil {
				continue
			}

			log.Println("Renaming", file.Name(), "to", res)

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

	// extract the name and change positions
	fName := strings.Join(tmp[:len(tmp)-2], "")
	fNum, err := strconv.Atoi(tmp[len(tmp)-1])
	if err != nil {
		return "", fmt.Errorf("%s no match", filename)
	}

	return fmt.Sprintf("%s_%d.%s", fName, fNum, extn), nil
}
