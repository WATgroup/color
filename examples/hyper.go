package main

import (
	"os"
	"oss.w-a-t.group/console/color"
)

func main() {

	//color.Init()
	color.EnableColor()	// override

	// lb := color.New(color.FgHiBlue)
	lnkCode := color.Hyperlink("https://www.w-a-t.group/", "test")
	// lb.Println("This is a ", lnkCode, " of hyperlinks")

	os.Stdout.WriteString(lnkCode)

	color.Red("\nbut this is regular text")

	os.Exit(0)
}
