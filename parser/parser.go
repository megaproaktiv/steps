package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const begin = "#begin"
const end = "#end"
const ignore = "#ignore"

// Parse
//
//	contentfile : filename
//	targetfile  : output filename
//	part: which part should be included, other parts will be
//	excluded
func Parse(contentfile *string, targetfile *string, part int) (int, error) {
	content, err := os.Open(*contentfile)
	if err != nil {
		log.Printf("Cannot read file %v\n", &contentfile)
		return 0, err
	}
	defer content.Close()

	fileScanner := bufio.NewScanner(content)

	fileScanner.Split(bufio.ScanLines)

	target, err := os.OpenFile(*targetfile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("Cannot create file %v\n", &targetfile)
		return 0, err
	}
	defer target.Close()
	// Scan 1 -> count
	count := 0
	insideBlock := false
	skipLine := false
	for fileScanner.Scan() {
		line := fmt.Sprintln(fileScanner.Text())
		if skipLine {
			skipLine = false
			continue
		}
		if IsIgnore(&line) {
			skipLine = true
			continue
		}
		if IsBegin(&line) {
			count += 1
			insideBlock = true
			continue
		}
		if IsEnd(&line) {
			insideBlock = false
			continue
		}
		if insideBlock {
			if count == part {
				_, err := target.Write([]byte(line))
				if err != nil {
					log.Printf("Cannot write line %v in  file %v\n",
						line,
						&targetfile)
					return 0, err
				}
				continue
			}
		} else {
			_, err := target.Write([]byte(line))
			if err != nil {
				log.Printf("Cannot write line %v in  file %v\n",
					line,
					&targetfile)
				return 0, err
			}
		}

	}

	return count, nil
}

func ParseTextOutput(contentfile *string, part int) (int, error) {
	content, err := os.Open(*contentfile)
	if err != nil {
		log.Printf("Cannot read file %v\n", &contentfile)
		return 0, err
	}
	defer content.Close()

	fileScanner := bufio.NewScanner(content)

	fileScanner.Split(bufio.ScanLines)

	// Scan 1 -> count
	count := 0
	insideBlock := false
	for fileScanner.Scan() {
		line := fmt.Sprintln(fileScanner.Text())
		if IsBegin(&line) {
			count += 1
			insideBlock = true
			continue
		}
		if IsEnd(&line) {
			insideBlock = false
			continue
		}
		if insideBlock {
			if count <= part {
				fmt.Print(line)

				continue
			}
		} else {

			fmt.Print(line)
		}

	}

	return count, nil
}

func IsBegin(line *string) bool {
	return strings.HasPrefix(*line, begin)
}
func IsEnd(line *string) bool {
	return strings.HasPrefix(*line, end)
}

func IsIgnore(line *string) bool {
	return strings.HasPrefix(*line, ignore)
}
