// SPDX-FileCopyrightText: Copyright (c) 2013 Fatih Arslan
// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: MIT

package color

import "io"

// Black is a convenient helper function to print with black foreground. A
// newline is appended to format by default.
func Black(format string, a ...any) { colorPrint(format, FgBlack, a...) }

// Red is a convenient helper function to print with red foreground. A
// newline is appended to format by default.
func Red(format string, a ...any) { colorPrint(format, FgRed, a...) }

// Green is a convenient helper function to print with green foreground. A
// newline is appended to format by default.
func Green(format string, a ...any) { colorPrint(format, FgGreen, a...) }

// Yellow is a convenient helper function to print with yellow foreground.
// A newline is appended to format by default.
func Yellow(format string, a ...any) { colorPrint(format, FgYellow, a...) }

// Blue is a convenient helper function to print with blue foreground. A
// newline is appended to format by default.
func Blue(format string, a ...any) { colorPrint(format, FgBlue, a...) }

// Magenta is a convenient helper function to print with magenta foreground.
// A newline is appended to format by default.
func Magenta(format string, a ...any) { colorPrint(format, FgMagenta, a...) }

// Cyan is a convenient helper function to print with cyan foreground. A
// newline is appended to format by default.
func Cyan(format string, a ...any) { colorPrint(format, FgCyan, a...) }

// White is a convenient helper function to print with white foreground. A
// newline is appended to format by default.
func White(format string, a ...any) { colorPrint(format, FgWhite, a...) }

////////////////////////////////////////////////////////////

// HiBlack is a convenient helper function to print with hi-intensity black foreground. A
// newline is appended to format by default.
func HiBlack(format string, a ...any) { colorPrint(format, FgHiBlack, a...) }

// HiRed is a convenient helper function to print with hi-intensity red foreground. A
// newline is appended to format by default.
func HiRed(format string, a ...any) { colorPrint(format, FgHiRed, a...) }

// HiGreen is a convenient helper function to print with hi-intensity green foreground. A
// newline is appended to format by default.
func HiGreen(format string, a ...any) { colorPrint(format, FgHiGreen, a...) }

// HiYellow is a convenient helper function to print with hi-intensity yellow foreground.
// A newline is appended to format by default.
func HiYellow(format string, a ...any) { colorPrint(format, FgHiYellow, a...) }

// HiBlue is a convenient helper function to print with hi-intensity blue foreground. A
// newline is appended to format by default.
func HiBlue(format string, a ...any) { colorPrint(format, FgHiBlue, a...) }

// HiMagenta is a convenient helper function to print with hi-intensity magenta foreground.
// A newline is appended to format by default.
func HiMagenta(format string, a ...any) { colorPrint(format, FgHiMagenta, a...) }

// HiCyan is a convenient helper function to print with hi-intensity cyan foreground. A
// newline is appended to format by default.
func HiCyan(format string, a ...any) { colorPrint(format, FgHiCyan, a...) }

// HiWhite is a convenient helper function to print with hi-intensity white foreground. A
// newline is appended to format by default.
func HiWhite(format string, a ...any) { colorPrint(format, FgHiWhite, a...) }

///////////////////////////////////////////////////////////////////////////////

// FprintFunc returns a new function that prints the passed arguments as
// colorized with color.Fprint().
func (c *Color) FprintFunc() func(w io.Writer, a ...any) {
	return func(w io.Writer, a ...any) {
		c.Fprint(w, a...)
	}
}

// PrintFunc returns a new function that prints the passed arguments as
// colorized with color.Print().
func (c *Color) PrintFunc() func(a ...any) {
	return func(a ...any) {
		c.Print(a...)
	}
}

// FprintfFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintf().
func (c *Color) FprintfFunc() func(w io.Writer, format string, a ...any) {
	return func(w io.Writer, format string, a ...any) {
		c.Fprintf(w, format, a...)
	}
}

// PrintfFunc returns a new function that prints the passed arguments as
// colorized with color.Printf().
func (c *Color) PrintfFunc() func(format string, a ...any) {
	return func(format string, a ...any) {
		c.Printf(format, a...)
	}
}

// FprintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Fprintln().
func (c *Color) FprintlnFunc() func(w io.Writer, a ...any) {
	return func(w io.Writer, a ...any) {
		c.Fprintln(w, a...)
	}
}

// PrintlnFunc returns a new function that prints the passed arguments as
// colorized with color.Println().
func (c *Color) PrintlnFunc() func(a ...any) {
	return func(a ...any) {
		c.Println(a...)
	}
}

// SprintFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprint(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output, example:
//
//	put := New(FgYellow).SprintFunc()
//	fmt.Fprintf(color.Output, "This is a %s", put("warning"))
func (c *Color) SprintFunc() func(a ...any) string {
	return func(a ...any) string {
		return c.Sprint(a...)
	}
}

// SprintfFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color) SprintfFunc() func(format string, a ...any) string {
	return func(format string, a ...any) string {
		return c.Sprintf(format, a...)
	}
}

// SprintlnFunc returns a new function that returns colorized strings for the
// given arguments with fmt.Sprintln(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.
func (c *Color) SprintlnFunc() func(a ...any) string {
	return func(a ...any) string {
		return c.Sprintln(a...)
	}
}

/*****

// BlackString is a convenient helper function to return a string with black
// foreground.
func BlackString(format string, a ...any) string { return colorString(format, FgBlack, a...) }

// RedString is a convenient helper function to return a string with red
// foreground.
func RedString(format string, a ...any) string { return colorString(format, FgRed, a...) }

// GreenString is a convenient helper function to return a string with green
// foreground.
func GreenString(format string, a ...any) string { return colorString(format, FgGreen, a...) }

// YellowString is a convenient helper function to return a string with yellow
// foreground.
func YellowString(format string, a ...any) string { return colorString(format, FgYellow, a...) }

// BlueString is a convenient helper function to return a string with blue
// foreground.
func BlueString(format string, a ...any) string { return colorString(format, FgBlue, a...) }

// MagentaString is a convenient helper function to return a string with magenta
// foreground.
func MagentaString(format string, a ...any) string {
	return colorString(format, FgMagenta, a...)
}

// CyanString is a convenient helper function to return a string with cyan
// foreground.
func CyanString(format string, a ...any) string { return colorString(format, FgCyan, a...) }

// WhiteString is a convenient helper function to return a string with white
// foreground.
func WhiteString(format string, a ...any) string { return colorString(format, FgWhite, a...) }

// HiBlack is a convenient helper function to print with hi-intensity black foreground. A
// newline is appended to format by default.
func HiBlack(format string, a ...any) { colorPrint(format, FgHiBlack, a...) }

// HiRed is a convenient helper function to print with hi-intensity red foreground. A
// newline is appended to format by default.
func HiRed(format string, a ...any) { colorPrint(format, FgHiRed, a...) }

// HiGreen is a convenient helper function to print with hi-intensity green foreground. A
// newline is appended to format by default.
func HiGreen(format string, a ...any) { colorPrint(format, FgHiGreen, a...) }

// HiYellow is a convenient helper function to print with hi-intensity yellow foreground.
// A newline is appended to format by default.
func HiYellow(format string, a ...any) { colorPrint(format, FgHiYellow, a...) }

// HiBlue is a convenient helper function to print with hi-intensity blue foreground. A
// newline is appended to format by default.
func HiBlue(format string, a ...any) { colorPrint(format, FgHiBlue, a...) }

// HiMagenta is a convenient helper function to print with hi-intensity magenta foreground.
// A newline is appended to format by default.
func HiMagenta(format string, a ...any) { colorPrint(format, FgHiMagenta, a...) }

// HiCyan is a convenient helper function to print with hi-intensity cyan foreground. A
// newline is appended to format by default.
func HiCyan(format string, a ...any) { colorPrint(format, FgHiCyan, a...) }

// HiWhite is a convenient helper function to print with hi-intensity white foreground. A
// newline is appended to format by default.
func HiWhite(format string, a ...any) { colorPrint(format, FgHiWhite, a...) }

// HiBlackString is a convenient helper function to return a string with hi-intensity black
// foreground.
func HiBlackString(format string, a ...any) string {
	return colorString(format, FgHiBlack, a...)
}

// HiRedString is a convenient helper function to return a string with hi-intensity red
// foreground.
func HiRedString(format string, a ...any) string { return colorString(format, FgHiRed, a...) }

// HiGreenString is a convenient helper function to return a string with hi-intensity green
// foreground.
func HiGreenString(format string, a ...any) string {
	return colorString(format, FgHiGreen, a...)
}

// HiYellowString is a convenient helper function to return a string with hi-intensity yellow
// foreground.
func HiYellowString(format string, a ...any) string {
	return colorString(format, FgHiYellow, a...)
}

// HiBlueString is a convenient helper function to return a string with hi-intensity blue
// foreground.
func HiBlueString(format string, a ...any) string { return colorString(format, FgHiBlue, a...) }

// HiMagentaString is a convenient helper function to return a string with hi-intensity magenta
// foreground.
func HiMagentaString(format string, a ...any) string {
	return colorString(format, FgHiMagenta, a...)
}

// HiCyanString is a convenient helper function to return a string with hi-intensity cyan
// foreground.
func HiCyanString(format string, a ...any) string { return colorString(format, FgHiCyan, a...) }

// HiWhiteString is a convenient helper function to return a string with hi-intensity white
// foreground.
func HiWhiteString(format string, a ...any) string {
	return colorString(format, FgHiWhite, a...)
}

***/

///////////////////////////////////////////////////////////////////////////////

// Equals returns a boolean value indicating whether two colors are equal.
func Equals(c1, c2 *Color) bool {
	if nil == c1 && nil == c2 {
		return true
	}
	if nil == c1 || nil == c2 {
		return false
	}

	if len(c1.attr) != len(c2.attr) {
		return false
	}

	counts := make(map[ColorAttrib]uint, len(c1.attr)+1)
	for _, a := range c1.attr {
		counts[a]++
	}

	for _, a := range c2.attr {
		if 0 == counts[a] {
			return false
		}
		counts[a]--
	}
	return true
}
