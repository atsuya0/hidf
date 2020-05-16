package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func init() {
	log.SetFlags(log.Ltime | log.Llongfile)
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("wrong number of arguments.")
		return
	}
	filePath := os.Args[1]

	fileForReading, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := fileForReading.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	file := file{
		path:           filepath.Dir(filePath),
		fileForReading: fileForReading,
	}

	if bool, err := file.isHiddenFile(); err != nil {
		log.Fatalln(err)
	} else if bool {
		fmt.Println("The hidden file is not supported.")
		return
	}

	tmp, err := os.Create(file.generateRandFilePath())
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := tmp.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	file.fileForWriting = tmp

	isHidden, err := file.isHidden()
	if err != nil {
		log.Fatal(err)
	}
	if isHidden {
		if err := file.extract(); err != nil {
			log.Fatal(err)
		}
		return
	}
	if err := file.hide(); err != nil {
		log.Fatal(err)
	}
}
