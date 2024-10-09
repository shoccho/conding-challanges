package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	flags := map[string]string{"-c": "bytes", "-l": "lines"}
	fmt.Println(process(args, flags))
}

func process(args []string, flags map[string]string) string {

	var output string
	if value, ok := flags[args[0]]; ok {
		file := processFile(args[1])
		switch value {
		case "bytes":
			output += fmt.Sprintf("%d ", len(file))
		case "lines":
			count := 0
			for _, b := range file {
				if b == '\n' {
					count++
				}
			}
			output += fmt.Sprintf("%d ", count)
		}
	}
	return output + args[1]
}

func processFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	return file

}
