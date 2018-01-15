package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func list(b bytes.Buffer) {
	i := 1
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			break
		}
		fmt.Printf("\t%v\t%v", i, string(line))
		i++
	}
}

func del(b bytes.Buffer, f *os.File, args []string) error {
	f.Truncate(0)
	f.Seek(0, 0)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if strings.HasPrefix(string(line), args[0]) {
			continue
		}
		f.Write(line)
	}
	return nil
}

func add(f *os.File, args []string) error {
	s, err := f.Stat()
	if err != nil {
		return err
	}
	b := []byte(strings.Join(args, " ") + "\n")
	f.WriteAt(b, s.Size())
	return nil
}

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
	if len(args) == 0 {
		list(b)
		return
	}
	if args[0] == "-d" {
		err = del(b, f, args[1:])
	} else {
		err = add(f, args[:])
	}
	err = f.Sync()
	if err != nil {
		panic(err)
	}
}
