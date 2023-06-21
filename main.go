package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	fileBoilerplate := getFileBoilerplate()

	err := ioutil.WriteFile(fileName, []byte(fileBoilerplate), 0644)
	if err != nil {
		fmt.Printf("%s '%s'\n",errorPrefix, err)
		return
	}

	fmt.Printf("Your file '%s' has been created. Ya welcome!\n", fileName)
}


func getFileBoilerplate() string {
	lineSeparator := "\n------------------------------\n\n\n"
	heading := "Journal Entry"
	howDidYouFeel := "How did you feel today?"
	whatDidYouDo := "What did you do today?"
	interestingThoughts := "Any interesting thoughts?"

	fileLines := []string{heading, howDidYouFeel, whatDidYouDo, interestingThoughts}


	return strings.Join(fileLines, lineSeparator) + lineSeparator
}