package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	flags := map[string]string{"-c": "bytes"}
	process(args, flags)
}

func process(args []string, flags map[string]string) {
	if value, ok := flags[args[0]]; ok {
		file := processFile(args[1])
		switch value {
		case "bytes":
			{
				fmt.Printf("%d %s\n", len(file), args[1])
			}

		}
	}
}

func processFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	return file

}
