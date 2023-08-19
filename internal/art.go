package internal

import (
	"ascii-art-web-dockerize/internal/files"
	"errors"
	"fmt"
	"strings"
)

func GenerateASCII(inputString string, style string) (string, error) {
	var filename string
	switch style {
	case "Standard":
		filename = "./internal/files/standard.txt"
	case "Shadow":
		filename = "./internal/files/shadow.txt"
	case "Thinkertoy":
		filename = "./internal/files/thinkertoy.txt"
	}
	splitSlice, err := files.ReadFile(filename)
	if err != nil {
		return "", errors.New("empty File")
	}
	result, errString := generator(inputString, splitSlice)
	if errString != nil {
		return "", errString
	}
	return result, nil
}

func generator(inputString string, splitSlice []string) (string, error) {
	ASCII := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
	runesASCII := []rune(ASCII)
	runesInputString := []rune(inputString)
	index := 0
	indexOfLetters := []int{}
	count := 0
	for index < len(runesInputString) {
		i := 0
		for i < len(ASCII) && index < len(runesInputString) {
			if index+1 < len(runesInputString) &&
				runesInputString[index] == '\\' &&
				runesInputString[index+1] == 'n' {
				j := index
				countOfSlashes := 0
				for j >= 0 && runesInputString[j] == '\\' {
					countOfSlashes++
					j--
				}
				if countOfSlashes%2 == 0 {
					indexOfLetters = append(indexOfLetters, 703)
				} else {
					indexOfLetters = append(indexOfLetters, -1)
				}
				i--
				index += 2
			} else if runesInputString[index] == runesASCII[i] {
				indexOfLetters = append(indexOfLetters, 9*i+1)
			} else if runesInputString[index] == rune(10) {
				indexOfLetters = append(indexOfLetters, -1)
				index++
			}
			i++
		}
		index++
	}
	splitSentenses := make([][]int, 0)
	startIndex := 0
	for i := 0; i < len(indexOfLetters); i++ {
		if indexOfLetters[i] == -1 {
			subSlice := indexOfLetters[startIndex:i]
			splitSentenses = append(splitSentenses, subSlice)
			startIndex = i + 1
			if i == len(indexOfLetters)-1 && !isSliceEmpty(splitSentenses) {
				splitSentenses = append(splitSentenses, []int{})
				startIndex = i + 1
			}
		}
	}
	if startIndex < len(indexOfLetters) {
		subSlice := indexOfLetters[startIndex:]
		splitSentenses = append(splitSentenses, subSlice)
	}
	var sbuild strings.Builder
	for j := 0; j < len(splitSentenses); j++ {
		count = 0
		for count != 8 && len(splitSentenses[j]) != 0 {
			i := 0
			for i < len(splitSentenses[j]) {
				if splitSentenses[j][i]+count < len(splitSlice) {
					fmt.Print(splitSlice[splitSentenses[j][i]+count])
					sbuild.WriteString(splitSlice[splitSentenses[j][i]+count])
				}
				i++
			}
			fmt.Println()
			sbuild.WriteString("\n")
			count++
		}
		if len(splitSentenses[j]) == 0 {
			fmt.Println()
			sbuild.WriteString("\n")
		}
	}
	return sbuild.String(), nil
}
func isSliceEmpty(splitSentenses [][]int) bool {
	for i := 0; i < len(splitSentenses); i++ {
		if len(splitSentenses[i]) != 0 {
			return false
		}
	}
	return true
}
