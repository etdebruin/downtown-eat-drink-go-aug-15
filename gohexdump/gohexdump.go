package gohexdump

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

type FileGrid struct {
	current  int64
	hexLines [][]byte
}

func (f FileGrid) Seek(offset int64, whence int) (n int64, err error) {

	f.current = offset

	return n, err
}

func (f FileGrid) Read(p []byte) (n int, err error) {

	p = f.hexLines[f.current]

	return n, err
}

func Dump(r io.Reader) io.ReadSeeker {

	var fileGrid FileGrid

	br := bufio.NewReader(r)

	for {
		hexLine := make([]byte, 16)
		_, err := br.Read(hexLine)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		fileGrid.hexLines = append(fileGrid.hexLines, hexLine)
	}

	for k, v := range fileGrid.hexLines {
		fmt.Printf("%d %s", k, hex.Dump(v))
	}

	return fileGrid

}
