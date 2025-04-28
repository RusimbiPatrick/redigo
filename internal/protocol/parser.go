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
func parseInLine(line string) (string, []string, error) {
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", nil, errors.New("empty command")
	}
	return strings.ToUpper(parts[0]), parts[1:], nil
}

func parseArray(reader *bufio.Reader, header string) (string, []string, error) {
    count, err := strconv.Atoi(strings.TrimSpace(header[1:]))
    if err != nil {
        return "", nil, err
    }
    var args []string
    for i := 0; i < count; i++ {
        line, err := reader.ReadString('\n')
        if err != nil {
            return "", nil, err
        }
        line = strings.TrimSpace(line)
        if len(line) > 0 && line[0] == '$' {
            length, err := strconv.Atoi(line[1:])
            if err != nil {
                return "", nil, err
            }
            bytes := make([]byte, length)
            _, err = reader.Read(bytes)
            if err != nil {
                return "", nil, err
            }
            args = append(args, string(bytes))
            _, err = reader.ReadString('\n') // Read CRLF
            if err != nil {
                return "", nil, err
            }
        } else {
            return "", nil, errors.New("invalid array format")
        }
    }
    if len(args) == 0 {
        return "", nil, errors.New("empty command")
    }
    return strings.ToUpper(args[0]), args[1:], nil
}