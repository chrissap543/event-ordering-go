package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	parse_input()
}

func parse_input() {
	args := os.Args[1:]
	identifier := args[0]
	input_filepath := args[1]

	input_file, err := os.Open(input_filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer input_file.Close()

	scanner := bufio.NewScanner(input_file)
	if !scanner.Scan() {
		log.Fatal("Input file empty")
	}

	num_peers, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("File does not start with number: ", err)
	}

	for i := 0; i < num_peers; i++ {
		if scanner.Scan() {
			text := scanner.Text()
			split := strings.Fields(text)

			if split[0] == identifier {
				port := split[2]
				continue
			}

		} else {
			log.Fatal("File does not contain all nodes")
		}
	}
}
