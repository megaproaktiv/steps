package main

import (
	"fmt"
	"log"
	"os"
	"steps/parser"
	"strconv"

	"github.com/urfave/cli/v2" // imports as package "cli"
)

// Run parser on given file
func main(){
	app := &cli.App{
        Name:  "parse",
        Usage: "parse a file for # begin:1 loops",
        Action: func(cCtx *cli.Context) error {
            fmt.Println("Parsing")
			input :=  cCtx.Args().Get(0)
			output :=  cCtx.Args().Get(1)
			countStr :=  cCtx.Args().Get(2)
			count, err := strconv.Atoi(countStr)
			if err != nil {
				log.Println("3rd argument is no integer")
				fmt.Println(err)
				os.Exit(2)
			}
			
			parser.Parse(&input, &output, count)
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}