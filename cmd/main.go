package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	isRecursive := flag.Bool("r", false, "Recursively rename files in subdirectories")
	flag.Parse()

	dirName := ""
	if len(os.Args) == 1 {
		fmt.Println("Please provide a directory name")
		return
	} else if len(os.Args) > 1 {
		dirName = os.Args[len(os.Args)-1]
	}
	if dirName == "" || dirName == "-r" {
		dirName = "."
	}

	fmt.Println("Reading files from", dirName, "recursively:", *isRecursive)

	if *isRecursive {
		err := filepath.WalkDir(dirName, func(path string, file os.DirEntry, err error) error {
			if err != nil {
				log.Fatal(err)
			}

			if file.IsDir() {
				return nil
			}

			res, err := matchAndConvert(file.Name())
			if err != nil {
				return nil
			}

			fmt.Printf("Renaming \"%s\" to \"%s\" \n", file.Name(), res)

			tmp := strings.Split(path, "/")
			tmp[len(tmp)-1] = res

			err = os.Rename(path, filepath.Join(tmp...))
			if err != nil {
				log.Fatal(err)
			}

			return nil

		})
		if err != nil {
			log.Fatal(err)
		}
	} else {
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

	fmt.Println("Done")

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
