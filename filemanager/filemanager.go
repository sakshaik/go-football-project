package filemanager

import (
	"bufio"
	"bytes"
	"os"
)

func ReadFileLines(path string) ([]string, error) {
	readFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines, nil
}

func ReadFileAsString(path string) (string, error) {
	readFile, err := os.Open(path)

	if err != nil {
		return "", err
	}
	defer readFile.Close()
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(readFile)
	data := buffer.String()
	return data, nil
}
