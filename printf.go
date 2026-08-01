// SPDX-FileCopyrightText: © 2025,2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

// Inspired + some code from fatih's "color" module

package color

import (
	"fmt"
	"io"
	"strings"
)

//revive:disable:indent-error-flow   False positive....
///////////////////////////////////////////////////////////////////////////////

func colorPrint(format string, p ColorAttrib, a ...any) {
	c := getCachedColor(p)

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	var sb1, sb2 strings.Builder

	c.sequence(&sb1)
	fmt.Fprint(stdOut, sb1.String())

	fmt.Fprintf(stdOut, format, a...)

	c.unseq(&sb2)
	fmt.Fprint(stdOut, sb2.String())
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

func (c Color) Print(a ...any) (n int, err error) {
	// var sb1, sb2 strings.Builder

	// escPrefix(&sb1)
	// c.sequence(&sb1)
	// escPostfix(&sb1)
	// fmt.Fprint(stdOut, sb1.String())

	// n, err = fmt.Fprint(stdOut, a...)

	// escPrefix(&sb2)
	// c.unseq(&sb2)
	// escPostfix(&sb2)
	// fmt.Fprint(stdOut, sb1.String())

	return c.Fprint(stdOut, a...)
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (c Color) Fprint(w io.Writer, a ...any) (n int, err error) {

	var sb1, sb2 strings.Builder

	escPrefix(&sb1)
	c.sequence(&sb1)
	escPostfix(&sb1)
	if n1, e1 := fmt.Fprint(w, sb1.String()); nil != e1 {
		return -1, e1
	} else {
		n += n1
	}

	nn, err := fmt.Fprint(w, a...)
	n += nn
	if nil != err {
		return
	}

	escPrefix(&sb2)
	c.unseq(&sb2)
	escPostfix(&sb2)
	if n2, e2 := fmt.Fprint(w, sb2.String()); nil != e2 {
		return -1, e2
	} else {
		n += n2
	}

	return
}

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func (c *Color) Fprintf(w io.Writer, format string, a ...any) (n int, err error) {
	var sb1, sb2 strings.Builder

	escPrefix(&sb1)
	c.sequence(&sb1)
	escPostfix(&sb1)
	if n1, e1 := fmt.Fprint(w, sb1.String()); nil != e1 {
		return -1, e1
	} else {
		n += n1
	}

	nn, err := fmt.Fprintf(w, format, a...)
	n += nn
	if nil != err {
		return
	}

	escPrefix(&sb2)
	c.unseq(&sb2)
	escPostfix(&sb2)
	if n2, e2 := fmt.Fprint(w, sb2.String()); nil != e2 {
		return -1, e2
	} else {
		n += n2
	}

	return
}

func (c Color) Printf(format string, a ...any) (n int, err error) {

	return c.Fprintf(stdOut, format, a...)
}
