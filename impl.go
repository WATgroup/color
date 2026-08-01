// SPDX-FileCopyrightText: © 2025,2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

// Inspired + some code from fatih's "color" module

package color

import (
	"io"
	"os"
	"strings"

	isatty "github.com/mattn/go-isatty"
)

const escape = "\x1b"
const escRune = 0x1b


func init() {
	Init()	// avoid crashes if/when user forgets to initialize the library
}



// checkNoColor returns true if $NO_COLOR is set to a non-empty string.
func checkNoColor() bool {
	return "" != os.Getenv("NO_COLOR")
}

// stdoutIsTerminal returns true if os.Stdout is a terminal.
// Returns false if os.Stdout is nil (e.g., when running as a Windows service).
func stdoutIsTerminal() bool {
	if os.Stdout == nil {
		return false
	}
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}

// SetDefaults sets stdOut/stdErr as needed
// they can be "nil" if running with no console (e.g. Windows service / detached daemon)
func SetDefaults() {
	stdOut = os.Stdout
	if nil == stdOut {
		stdOut = io.Discard
	}

	stdErr = os.Stderr
	if nil == stdErr {
		stdErr = io.Discard
	}
}

func fputbs(w io.Writer, x []byte) {
 	w.Write(x)
}
func fputs(w io.Writer, s string) {
	w.Write([]byte(s))
}

///////////////////////////////////////////////////////////////////////////////

// sequence returns a formatted SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan

//go:noinline
func (c *Color) sequence(sb *strings.Builder) {
	for i, v := range c.attr {
		if i > 0 {
			sb.WriteByte(';')
		}
		sb.WriteString(smallNum(uint16(v)))
	}
}

// wrap wraps the s string with the colors attributes.
// The string is ready to be printed.
//
//go:noinline
func (c *Color) wrap(s string) string {
	if disableColor {
		return s
	}
	var sb strings.Builder
	// return c.format() + s + c.unformat()

	// PREFIX:   ESC[X..m
	escPrefix(&sb)
	c.sequence(&sb)
	escPostfix(&sb)

	// insert the string
	sb.WriteString(s)

	// POSTFIX:  ESC[X..m
	escPrefix(&sb)
	c.unseq(&sb)
	escPostfix(&sb)

	return sb.String()
}

//go:noinline
func (c *Color) unseq(sb *strings.Builder) {
	// for each element in sequence:
	// - use the specific reset escape if available,
	// - or the generic one if not found
	rr := smallNum(uint16(Reset))
	for i, v := range c.attr {
		if i > 0 {
			sb.WriteByte(';')
		}
		if ra, ok := mapResetAttributes[v]; ok {
			sb.WriteString(smallNum(uint16(ra)))
		} else {
			sb.WriteString(rr)
		}
	}
}

func (c *Color) format() string {
	var sb strings.Builder

	escPrefix(&sb)
	c.sequence(&sb)
	escPostfix(&sb)

	return sb.String()
}

func escPrefix(sb *strings.Builder) {
	sb.WriteRune(escRune)
	sb.WriteRune('[')
}
func escPostfix(sb *strings.Builder) {
	sb.WriteRune('m')
}
