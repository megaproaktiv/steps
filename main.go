package main

import (
	
	"fmt"
	"os"
	"steps/parser"

	"github.com/alecthomas/kong"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var CLI struct {
	Step struct {
		PathsIn string `arg:"" name:"pathin" help:"Paths to text file." type:"path"`
		PathsOut string `arg:"" name:"pathout" help:"Paths to text file." type:"path"`
		Count int `arg:"" name:"max" help:"maximum blocks" type:"int"`
	} `cmd:"" help:"render #begin #end blocks increasing and decreasing."`
}

// Run parser on given file
func main() {
	var d2Path string
	var outfile string
	var max int
	// Open the PowerPoint file and parse the XML data
	ctx := kong.Parse(&CLI,
		kong.Name("stepper"),
		kong.Description("render #begin #end blocks increasing and decreasing."),
	)
	switch ctx.Command() {
	case "step <pathin> <pathout> <max>":
		d2Path = CLI.Step.PathsIn
		outfile = CLI.Step.PathsOut
		max = CLI.Step.Count
		fmt.Println(d2Path)
	default:
		panic(ctx.Command())
	}
	// for {
		// 	parser.ParseTextOutput(&d2Path, number)
		
		// 	reader := bufio.NewReader(os.Stdin)
		// 	char, err := reader.ReadByte()
		// 	if err != nil {
			// 		fmt.Println(err)
			// 	}
			
			// 	if char == 27 { // ESC
				// 		break
				// 	} else if char == 67 { // Cursor right
					// 		number++
					// 	} else if char == 68 { // Cursor left
						// 		number--
						// 	}
						
						
						// }
						
	number := 1
	fmt.Println(number, " ", "+ - with left and right, stop with esc")
	parser.Parse(&d2Path, &outfile, number)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		fmt.Println(number)
		switch key.Code {
		case keys.CtrlC, keys.Escape:
		  return true, nil // Return true to stop listener
		case keys.Right:
			number++
			if number > max {
				number = max
			}
			parser.Parse(&d2Path, &outfile, number)
		case keys.Left:
			number--
			if number == 0 {
				number = 1
			}
			parser.Parse(&d2Path,&outfile, number)
		case keys.RuneKey: // Check if key is a rune key (a, b, c, 1, 2, 3, ...)
		  if key.String() == "q" { // Check if key is "q"
			fmt.Println("\rQuitting application")
			os.Exit(0) // Exit application
		  }
		  fmt.Printf("\rYou pressed the rune key: %s\n", key)
		default:
		  fmt.Printf("\rYou pressed: %s\n", key)
		}
	  
		return false, nil // Return false to continue listening
	  })
}
