package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile(os.Getenv("TODO_FILE"), os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Could not open todo file. Is $TODO_FILE set?")
		return
	}
	defer f.Close()
	args := os.Args[1:]
	var b bytes.Buffer
	io.Copy(&b, f)
	switch len(args) {
	// print
	case 0:
		i := 1
		for {
			line, err := b.ReadBytes('\n')
			if err != nil {
				break
			}
			fmt.Printf("\t%v  %v", i, string(line))
			i++
		}
	// add
	case 1:
		s, err := f.Stat()
		if err != nil {
			panic(err)
		}
		b := []byte(args[0] + "\n")
		f.WriteAt(b, s.Size())
		fmt.Println("OK")
	// del
	case 2:
		f.Truncate(0)
		f.Seek(0, 0)
		for {
			line, err := b.ReadBytes('\n')
			if err != nil {
				break
			}
			if strings.HasPrefix(string(line), args[1]) {
				continue
			}
			f.Write(line)
		}
		fmt.Println("OK")
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}
}
