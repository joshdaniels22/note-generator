package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var errorPrefix string = "An error occurred:"

func main() {
	currentTime := time.Now()

	date := currentTime.Format("06-01-02")
	fileName := fmt.Sprintf("%s-Note.txt", date)

	if _, err := os.Stat(fileName); err == nil {
		fileExists()
	} else {
		fileDoesNotExist(fileName)
	}
}

func fileExists() {
	fmt.Printf("That file already exists. In future, please check before wasting my time :)\n")
}

func fileDoesNotExist(fileName string) {

	err := ioutil.WriteFile(fileName, []byte("Journal Entry\n\n"), 0644)
	if err != nil {
		fmt.Printf("%s '%s'\n",errorPrefix, err)
		return
	}

	fmt.Printf("Your file '%s' has been created. Ya welcome!\n", fileName)
}
