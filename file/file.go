package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type fileHandler struct {
}

func GetFileHandler() *fileHandler {
	fh := &fileHandler{}
	return fh
}

func (fh *fileHandler) ReadFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (fh *fileHandler) WriteFile(path string, content string) error {
	if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

func (fh *fileHandler) ReadDir(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	return files
}

func (fh *fileHandler) ReadLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)
	var result []string
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				return result, nil
			}
			return nil, err
		}
		result = append(result, line)
	}
}
