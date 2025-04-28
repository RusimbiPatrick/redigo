package protocol

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"strings"
)


func Parse(input []byte)(string, []string, error){
	reader := bufio.NewReader(bytes.NewReader(input))

	line, err := reader.ReadString('\n')
	if err != nil {
		return "", nil, err
	}
	if line[0] == '*' {
		return parseArray(reader, line)
	}
	return parseInLine(line)
}

func parseArray(reader *bufio.Reader, header string)(string, []string, error){
	cout, err := strconv.Atoi(header[1: len(header)-2])
	if err != nil {
		return "", nil, err
	}
	var args []string
	for i := 0; i < count; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", nil, err
		}
		if line[0] = '$' {
			length, _ := strconv.Atoi(line[1: len(line)-2])
			 bytes := make([]byte, length)
			 reader.Read(bytes)
			 args = append(args, string(bytes))
			 reader.ReadString('\n') //CRLF
		}
	}
	if len(args) == 0 {
		return "", nil, errors.New("empty command")
	}
	return strings.ToUpper(args[0]), args[1:], nil
}