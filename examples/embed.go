package main

import (
	"os"
	"oss.w-a-t.group/console/color"
)

func main() {

	//color.Init()
	color.EnableColor()	// override

	b := color.New(color.FgHiBlue)
	fn := b.FprintFn()

	cn := 3
	fn(os.Stdout,"This is blue","and","so is",cn,"\n")

	os.Exit(0)
}
