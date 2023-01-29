package parser

import (
	"bufio"
	"log"
	"os"
	"strings"
)


const delim = "#"
const begin = "# begin:"
const end = "# end"

// Parse
// 		contentfile : filename
// 		targetfile  : output filename
// 		part: which part should be included, other parts will be 
// 		excluded
func Parse(contentfile *string, targetfile *string, part int) (int,error){
	content,err := os.Open(*contentfile)
	if err != nil{
		log.Printf("Cannot read file %v\n", &contentfile)
		return 0,err
	}
	defer content.Close()
	
	fileScanner := bufio.NewScanner(content)
	
    fileScanner.Split(bufio.ScanLines)
	
	target ,err := os.OpenFile(*targetfile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Printf("Cannot create file %v\n", &targetfile)
		return 0,err
	}
	defer target.Close()
	// Scan 1 -> count
	count := 0
	include := false
    for fileScanner.Scan() {
		line := fileScanner.Text()
		if (count == part && include){
			_, err := target.Write([]byte(line))
			if err != nil {
				log.Printf("Cannot write line %v in  file %v\n",
				 line,
				 &targetfile)
				return 0,err
			}
		}
        if IsBegin(&line){
			count += 1
			include = true
		}
    }

	return count,nil
}

func IsBegin(line *string) bool{
	return strings.HasPrefix(*line,begin )
}