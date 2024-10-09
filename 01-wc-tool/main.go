package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type flags struct {
	countByte bool
	countChar bool
	countWord bool
	countLine bool
}

func main() {
	flags := processFlags()
	in, filename := processInput()
	if filename != "" {
		defer in.(*os.File).Close()
	}
	processOutput(in, flags)
}

func processFlags() flags {
	countByte := flag.Bool("c", false, "Count bytes")
	countChar := flag.Bool("m", false, "Count charecters")
	countWord := flag.Bool("w", false, "Count words")
	countLine := flag.Bool("l", false, "Count liness")

	flag.Prase()
	if !(*countByte || *countWord || *countLine || *countChar) {
		*countByte = true
		*countWord = true
		*countLine = true
	}

	return flags{*countByte, *countChar, *countWord, *countLine}

}
func processInput() (io.Reader, string) {
	filename := ""
	if filename = flag.Arg(0); filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		return file, filename
	}
	return os.Stdin, filename
}

func isWhiteSpace(b rune) bool {
	return b == ' ' || b == '\n' || b == '\t' || b == '\r'
}

func processOutput(in io.Reader, f flags) {
	buf := bufio.NewReader(in)
	lastByte := rune(0)
	byteCount, wordCount, charCount, lineCount := 0, 0, 0, 0
	for {
		curByte, sz, err := buf.ReadRune()
		if err == io.EOF {
			if !isWhiteSpace(lastByte) {
				wordCount++
			}
			break
		}
		byteCount += sz
		charCount++
		if isWhiteSpace(curByte) && !isWhiteSpace(lastByte) {
			wordCount++
		}
		if curByte == '\n' {
			lineCount++
		}
		lastByte = curByte
	}
	if f.countLine {
		fmt.Print(lineCount, " ")
	}
	if f.countChar {
		fmt.Print(charCount, " ")
	}
	if f.countWord {
		fmt.Print(wordCount, " ")
	}
	if f.countByte {
		fmt.Print(byteCount, " ")
	}
}
