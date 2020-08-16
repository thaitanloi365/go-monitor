package docker

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"strings"
)

func logReader(reader io.ReadCloser, tty bool) func() (string, error) {
	if tty {
		scanner := bufio.NewScanner(reader)
		return func() (string, error) {
			if scanner.Scan() {
				return scanner.Text(), nil
			}

			return "", io.EOF
		}
	}
	hdr := make([]byte, 8)
	var buffer bytes.Buffer
	return func() (string, error) {
		buffer.Reset()
		_, err := reader.Read(hdr)
		if err != nil {
			return "", err
		}
		count := binary.BigEndian.Uint32(hdr[4:])
		_, err = io.CopyN(&buffer, reader, int64(count))
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(buffer.String()), nil
	}
}
