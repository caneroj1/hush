package hush

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Hush struct {
	properties map[string]string
}

func Hushfile() Hush {
	// open the hushfile
	f := openHushfile()
	var h Hush
	scanner := bufio.NewScanner(f)
	h.properties = readHushFile(scanner)

	return h
}

func (h Hush) GetString(key string) (string, bool) {
	val, ok := h.properties[key]

	if !ok {
		return "", ok
	}
	return val, true
}

func (h Hush) GetInt(key string) (int64, bool) {
	val, ok := h.properties[key]
	var ret int64

	if !ok {
		return ret, ok
	}

	ret, err := strconv.ParseInt(val, 0, 64)
	if err != nil {
		return ret, false
	}
	return ret, true
}

func (h Hush) GetFloat(key string) (float64, bool) {
	val, ok := h.properties[key]
	var ret float64

	if !ok {
		return ret, ok
	}

	ret, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return ret, false
	}
	return ret, true
}

func openHushfile() *os.File {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// try and open the hushfile if it is in the same
	// directory
	f, err := os.Open(filepath.Join(path, ".hushfile"))
	if err != nil {
		// try and open the hushfile if it is in /conf/
		f, err = os.Open(filepath.Join(path, "/conf/.hushfile"))
		if err != nil {
			// try and open the hushfile if it is in /app/
			// this is the case with a Revel application
			f, err = os.Open(filepath.Join(path, "/app/.hushfile"))
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
		}
	}

	return f
}

// read the hushfile line by line
func readHushFile(scanner *bufio.Scanner) map[string]string {
	properties := make(map[string]string)
	lineno := 1

	// read the file and process the lines, splitting them by ':'
	for scanner.Scan() {
		pair := processLine(scanner.Text())
		if len(pair) != 2 {
			log.Fatalf("Hushfile line %d was formatted improperly", lineno)
			panic("Hushfile Error!")
		}

		// add the hushfile properties and their values to the properties map
		properties[pair[0]] = pair[1]
		lineno++
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return properties
}

func processLine(line string) []string {
	// trim whitespace from the line
	tokens := strings.Split(strings.TrimSpace(line), ":")

	// trim whitespace from each token, if there is any
	for ind, val := range tokens {
		tokens[ind] = strings.TrimSpace(val)
	}

	return tokens
}
