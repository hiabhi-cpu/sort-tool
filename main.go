package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Hello")
	fileName, err := getFileName(os.Args)
	check(err)
	// fmt.Println(fileName)
	commands, err := getCommands(os.Args)
	check(err)
	readFile(fileName, commands)
}

func readFile(fileName string, commands []string) {
	data, err := os.ReadFile(fileName)
	check(err)

	lines := strings.Split(string(data), "\n")

	filterdLines := make([]string, 0)
	for _, r := range lines {
		r = strings.TrimSpace(r)
		if r != "" {
			filterdLines = append(filterdLines, r)
		}
	}

	// fmt.Println(len(lines))
	// slices.Sort(filterdLines)

	// temp := lines[0:10]
	// fmt.Println(temp)
	// temp = mergeSort(temp)
	// fmt.Println(temp)

	filterdLines = mergeSort(filterdLines)

	finalList, err := filterOnCommand(filterdLines, commands)
	check(err)

	for _, r := range finalList {
		fmt.Println(r)
	}

}

func filterOnCommand(filterdLines, commands []string) ([]string, error) {
	if len(commands) == 1 {
		return filterdLines, nil
	}
	var resultList []string
	flag := 0
	for _, r := range commands {
		if r == "-u" {
			flag = 1
			resultList = getUniqueLines(filterdLines)
		}
	}
	if flag == 0 {
		return nil, errors.New("Give correct option")
	}
	return resultList, nil
}

func getUniqueLines(lines []string) []string {
	unique := make([]string, 0)
	prev := ""
	for _, l := range lines {
		if l != prev {
			unique = append(unique, l)
			prev = l
		}
	}
	return unique
}

func getFileName(args []string) (string, error) {
	for _, r := range args {
		if strings.Contains(r, ".txt") {
			return r, nil
		}
	}
	return "", errors.New("No text file ive correct file name in the args")
}

func getCommands(args []string) ([]string, error) {
	result := make([]string, 0)
	for _, r := range args {
		if strings.Contains(r, "-") {
			result = append(result, r)
		}
	}
	return result, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
