package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func addFiles(filename string, zipW *zip.Writer) error {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	splitted_filename := strings.Split(filename, "/")
	filename_new := splitted_filename[len(splitted_filename)-1]

	wr, err := zipW.Create(filename_new)
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(wr, file); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC

	file, err := os.OpenFile("bonus2.zip", flags, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	if len(os.Args) == 2 {
		filenames := strings.Split(os.Args[1], ",")
		for _, filename := range filenames {
			if err := addFiles(filename, zipWriter); err != nil {
				panic(err)
			}
		}
	} else {
		for _, filename := range os.Args[1:] {
			if err := addFiles(filename, zipWriter); err != nil {
				panic(err)
			}
		}
	}
}
