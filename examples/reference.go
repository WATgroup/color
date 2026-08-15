package main

import (
	"fmt"
	"github.com/WATgroup/errors"
	"os"
	"oss.w-a-t.group/console/color"
)

var (
	myWriter = os.Stderr
)

func main() {

	sample1()
	sample2()
	sample3()
	sample4()
	sample5()

	os.Exit(0)
}

func sample1() {

	color.Cyan("Prints text in cyan.")

	// a newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// More default foreground colors..
	color.Red("We have red")
	color.Yellow("Yellow color too!")
	color.Magenta("And many others ..")

	// Hi-intensity colors
	color.HiGreen("Bright green color.")
	color.HiBlack("Bright black means gray..")
	color.HiWhite("Shiny white color!")

}

func sample2() {

	// Create a new color object
	c := color.NewP(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with White background.")

	// Use your own io.Writer output
	color.New(color.FgBlue).Fprintln(myWriter, "blue color!")

	blue := color.New(color.FgBlue)
	blue.Fprint(myWriter, "This will print text in blue.")

}

func sample3() {
	err := errors.New("this is an error message")

	// Create a custom print function for convenient
	red := color.New(color.FgRed).PrintfFunc()
	red("warning")
	red("error: %s", err)

	// Mix up multiple attributes
	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice("don't forget this...")

}

func sample4() {
	stars := "We are all made of stars!"

	blue := color.New(color.FgBlue).FprintfFunc()
	blue(myWriter, "important notice: %s", stars)

	// Mix up with multiple attributes
	success := color.New(color.Bold, color.FgGreen).FprintlnFunc()
	success(myWriter, "don't forget this...")

}

func sample5() {
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	fmt.Printf("this is a %s and this is %s.\n", yellow("warning"), red("error"))

	info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
	fmt.Printf("this %s rocks!\n", info("package"))

}
