package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func fileHandler(fileName string) {

	cmdFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer cmdFile.Close()
	cmdScanner := bufio.NewScanner(cmdFile)
	for cmdScanner.Scan() {
		cmdInput := cmdScanner.Text()
		cmdInput = strings.TrimRight(cmdInput, "\n")
		if cmdInput != "" {
			output := Run(cmdInput)
			fmt.Println(output)
		}
	}

	if err := cmdScanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func Run(input string) string {
	fmt.Println(input)
	return ""
}
