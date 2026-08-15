package main

import (
	"os"
	"oss.w-a-t.group/console/color"
)

func main() {

	//color.Init()
	// color.EnableColor()	// override

	b := color.New(color.FgHiBlue)
	b.Println("This is blue")

	b.Add(color.Underline)
	b.Print("This is *underlined* blue\n")

	var c color.Color
	c.Add(color.FgRed, color.CrossedOut)

	c.Printf("%s\n", "an error message")

	os.Exit(0)
}
