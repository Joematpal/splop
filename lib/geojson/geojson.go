package geojson

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Iterator interface {
	Next() bool
	String() string
	Error() error
	Ch() chan string
	Close() error
}

type iterator struct {
	next   chan string
	err    chan error
	file   *os.File
	primer bool
}

func (it *iterator) Next() bool {
	return len(it.next) > 0 || it.primer
}

func (it *iterator) String() string {
	s := <-it.next
	return s
}
func (it *iterator) Error() error {
	return <-it.err
}
func (it *iterator) Ch() chan string {
	return it.next
}

func (it *iterator) Close() error {
	var err error
	defer func() {
		close(it.err)
	}()
	if it.file != nil {
		if err := it.file.Close(); err != nil {
			if len(it.err) != 0 {
				_ = <-it.err
			}
			return err
		}
	}
	if len(it.err) > 0 {
		err = <-it.err
	}

	return err
}

func ReadGeoJson(filePath string) (Iterator, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	itr := &iterator{
		next: make(chan string, 2),
		err:  make(chan error, 1),
		file: file,
	}

	reader := bufio.NewReader(file)
	i := -1
	go func() {
		for {
			i++
			line, rerr := reader.ReadString('\n')
			if rerr != nil {
				if rerr != io.EOF {
					itr.err <- rerr
				}
				defer close(itr.next)
				return
			}
			if i == 0 {
				continue
			}

			str := strings.Trim(strings.TrimSpace(line), ",")
			v := map[string]interface{}{}

			if err := json.Unmarshal([]byte(str), &v); err != nil {
				itr.err <- err
				if rerr != io.EOF {
					close(itr.next)
				}
				return
			}
			itr.next <- str
		}
	}()

	return itr, nil
}

func LoadGeoJson(url, filePath string) error {
	iter, err := ReadGeoJson(filePath)
	if err != nil {
		return err
	}

	for s := range iter.Ch() {
		resp, err := http.Post(url, "application/json", bytes.NewBufferString(s))
		if err != nil {
			fmt.Println("resp:", err)
			log.Print(err)
		} else {
			log.Print(resp)
		}
	}

	return iter.Close()
}
