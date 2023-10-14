package tcp

import (
	"bufio"
	"io"
)

const delim = '\n'

func ReadString(r io.Reader) (string, error) {
	str, err := bufio.NewReader(r).ReadString(delim)
	if err != nil {
		return "", err
	}
	return str[:len(str)-1], nil
}

func WriteString(w io.Writer, str string) error {
	_, err := w.Write(append([]byte(str), delim))
	return err
}
