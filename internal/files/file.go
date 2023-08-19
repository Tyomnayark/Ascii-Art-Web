package files

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadFile(filepath string) ([]string, error) {
	content, err := ioutil.ReadFile(filepath)
	hash := sha256.Sum256(content)
	hashString := fmt.Sprintf("%x", hash)
	if hashString != "c3ec7584fb7ecfbd739e6b3f6f63fd1fe557d2ae3e24f870730d9cf8b2559e94" {
		log.Println("Template file was changed")
	}
	if err != nil {
		log.Println(err)
	}
	var splitSlice = []string{}
	fileInput, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer fileInput.Close()
	scanner := bufio.NewScanner(fileInput)
	// create a slice where each line from the template is located separately
	for scanner.Scan() {
		line := scanner.Text()
		splitSlice = append(splitSlice, line)
	}
	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return splitSlice, nil
}
