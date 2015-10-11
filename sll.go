// sll ("strip long line") will remove long lines from a grep output.
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const DEFAULT_STRIP_SIZE = 1024

var help = `Usage: sll [size]

	size: Strip lines longer than this from the command output. Defaults to 1024
`

func getLength(args []string) (int, error) {
	argLen := len(args)
	if argLen > 0 {
		if argLen >= 2 {
			return 0, errors.New("too many arguments provided")
		}
		length, err := strconv.ParseUint(args[0], 10, 16)
		if err != nil {
			return 0, err
		}
		if length <= 0 {
			return 0, fmt.Errorf("invalid line length: %d", length)
		}
		return int(length), nil
	} else {
		return 1024, nil
	}
	return 0, nil
}

func init() {
	flag.Usage = func() {
		os.Stderr.Write([]byte(help))
	}
}

func stripLongLines(length int, in io.Reader, out io.Writer) error {
	buf := bufio.NewReaderSize(in, length)
	for {
		line, isPrefix, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if isPrefix {
			// line was too long. ignore it
			for isPrefix {
				_, isPrefix, err = buf.ReadLine()
				if err != nil {
					if err == io.EOF {
						break
					}
					log.Fatal(err)
				}
			}
		} else {
			out.Write(line)
			out.Write([]byte("\n"))
		}
	}
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	length, err := getLength(args)
	if err != nil {
		log.Fatal(err)
	}
	err = stripLongLines(length, os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
