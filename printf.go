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

func colorPrint(format string, p ColorAttrib, args ...any) {

	if disableColor {
		fputs(stdOut, fmt.Sprintf(format, args...))
		return
	}

	c := getCachedColor(p)

	if !strings.HasSuffix(format, "\n") { // ?? compat....
		format += "\n"
	}

	var sb strings.Builder

	c.sequence(&sb)
	n, _ := fmt.Fprintf(&sb, format, args...)
	c.unseq(&sb)
	fputsb(stdOut, &sb, n)
}

///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

// Print formats using the default formats for its operands and writes to
// standard output.
// Returns the number of bytes written and any write error encountered.
// This is the standard fmt.Print() method wrapped with the given color
func (c *Color) Print(args ...any) (n int, err error) {
	if disableColor {
		return fmt.Fprint(stdOut,args...)
	}

	var sb strings.Builder

	escPrefix(&sb)
	c.sequence(&sb)
	escPostfix(&sb)

	n, err = fmt.Fprint(&sb, args...)
	if nil != err {
		return -1, err
	}

	escPrefix(&sb)
	c.unseq(&sb)
	escPostfix(&sb)

	return fputsb(stdOut, &sb, n)
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (c *Color) Fprint(w io.Writer, args ...any) (n int, err error) {
	var sb strings.Builder
	n, err = c.wrapsbv(&sb, args)
	if nil != err {
		return
	}
	return fputsb(w, &sb, n)
}

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
func (c *Color) Fprintf(w io.Writer, format string, args ...any) (n int, err error) {
	var sb strings.Builder
	n, err = c.wrapsbf(&sb, format, args)
	if nil != err {
		return
	}
	return fputsb(w, &sb, n)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// This is the standard fmt.Printf() method wrapped with the given color.
func (c *Color) Printf(format string, args ...any) (n int, err error) {
	return c.Fprintf(stdOut, format, args...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.
func (c *Color) Println(args ...any) (n int, err error) {
	var sb strings.Builder
	n, err = c.wrapsbv(&sb, args)
	if nil != err {
		return
	}
	sb.WriteString(lineFeed)
	return fputsb(stdOut, &sb, n+1)
}

// Sprint is just like Print, but returns a string instead of printing it.
func (c *Color) Sprint(args ...any) string {
	// return c.wrap(fmt.Sprint(args...))
	var sb strings.Builder
	c.wrapsbv(&sb, args)
	return sb.String()
}

// Sprintln is just like Println, but returns a string instead of printing it.
func (c *Color) Sprintln(args ...any) string {
	var sb strings.Builder
	c.wrapsbv(&sb, args)
	sb.WriteString(lineFeed)
	return sb.String()
}

// Sprintf is just like Printf, but returns a string instead of printing it.
func (c *Color) Sprintf(format string, args ...any) string {
	var sb strings.Builder
	c.wrapsbf(&sb, format, args)
	return sb.String()
}

///////////////////////////////////////////////////////////////////////////////

//go:noinline
func (c *Color) wrapsbv(psb *strings.Builder, args []any) (n int, err error) {
	if disableColor {
		return fmt.Fprint(psb, args...)
	}

	// PREFIX:   ESC[X..m
	escPrefix(psb)
	c.sequence(psb)
	escPostfix(psb)

	// insert the string
	n, err = fmt.Fprint(psb, args...)

	// POSTFIX:  ESC[Y..m
	escPrefix(psb)
	c.unseq(psb)
	escPostfix(psb)

	return n, err
}

//go:noinline
func (c *Color) wrapsbf(psb *strings.Builder, format string, args []any) (n int, err error) {
	if disableColor {
		return fmt.Fprintf(psb, format, args...)
	}

	// PREFIX:   ESC[X..m
	escPrefix(psb)
	c.sequence(psb)
	escPostfix(psb)

	// insert the string
	n, err = fmt.Fprintf(psb, format, args...)

	// POSTFIX:  ESC[X..m
	escPrefix(psb)
	c.unseq(psb)
	escPostfix(psb)

	return n, err
}
