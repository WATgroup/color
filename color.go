// SPDX-FileCopyrightText: © 2025,2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

// Package color provides mid-level ANSI-sequence-based management primitives
// for (colored) console Output
package color


//** Implementation based upon github.com/fatih/color, distributed under MIT license
//   ...but otherwise completely reimplemented using modern patterns.
//   The original was from 2013 and it shows :O
//   Whenever feasible, this library's interface is the same as fatih/color
//   for compatibility sake. This version should be better in every aspect,
//   including performance, memory efficiency and correctness.
//**/

import (
	"fmt"
	"io"
)

var (
	stdOut       io.Writer
	stdErr       io.Writer
	disableColor = checkNoColor() || !stdoutIsTerminal()
)

//revive:disable-next-line:exported It's ok...
type ColorAttrib uint16

// Color represents a console "color", in terms of attributes/ESC-codes
type Color struct {
	attr []ColorAttrib
}

// Initialize the library -- optional
func Init() {
	disableColor = checkNoColor() || !stdoutIsTerminal()
	SetDefaults()
}

func SetOutput(w io.Writer) {
	stdOut = w
}

///////////////////////////////////////////////////////////////////////////////

// New returns a newly created color object.
func New(value ...ColorAttrib) (c Color) {
	if disableColor { // checkNoColor
		return
	}
	c.Add(value...)
	return
}

///////////////////////////////////////////////////////////////////////////////

// DisableColor disables colored output, returning the previous state
func DisableColor() (d bool) {
	d, disableColor = disableColor, true
	return
}

// EnableColor force-enables coloring (e.g. container console output / when not a TTY)
// and returns the previous value
func EnableColor() (d bool) {
	d, disableColor = disableColor, false
	return
}

// Unset resets terminal state ("ESC[0m")
func Unset() {
	// Should be called "Reset"... but conflicts with the ESC-reset name
	fmt.Fprint(stdOut, "%s[%dm", escape, Reset)
}

// Add includes new attributes into the color definition
func (c *Color) Add(value ...ColorAttrib) *Color {
	c.attr = append(c.attr, value...)
	return c
}

// Puts outputs a *colored* string to stdOut [no formatting/optimized]
func (c Color) Puts(s string) {
	fputs(stdOut, c.wrap(s))
}
