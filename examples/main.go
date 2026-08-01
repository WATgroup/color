
package main

import (
    "oss.w-a-t.group/console/color"
    "os"
)


func main() {

    //color.Init()

    // color.EnableColor()	// override

    b := color.New(color.FgHiBlue)
    b.Println("This is blue")

    b.Add(color.Underline)
    b.Print("This is *underlined* blue")

    os.Exit(0)
}
