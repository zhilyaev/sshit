package watcher

import (
	"bufio"
	"io"
	"os"
	"time"
)

func OnChange(filename string, f func(s string)) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	go Watcher(bufio.NewReader(file), f)
}

func Watcher(rd *bufio.Reader, f func(s string)) {
	for {
		time.Sleep(200 * time.Millisecond)
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			continue
		} else if err != nil {
			panic(err)
		}

		s := string(line)
		f(s)
	}
}
