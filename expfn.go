// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package color

// FprintFn returns a new function that prints the passed arguments colorized
// as with color.Fprint().  Escape sequences get cached at this point
func (c Color) FprintFn() func(w ioWriter, a ...any) (int, error) {
	if disableColor { // shortcut, if color disabled
		return fprintFn
	}

	cce := ccOp(c.attr[0])
	return func(w ioWriter, a ...any) (n int, err error) {
		var nr int
		nr, _ = w.Write([]byte(cce.set)) // should never fail
		n += nr
		if nr, err = fprintFn(w, a...); nil != err {
			return -1, err
		}
		n += nr
		nr, _ = w.Write([]byte(cce.unset))
		n += nr
		return
	}
}

// FputsFn returns a new function that prints the passed arguments colorized
// as with color.Puts().  Escape sequences get cached at this point
func (c Color) FputsFn() func(w ioWriter, s string) error {
	if disableColor { // shortcut, if color disabled
		return fputs
	}

	cce := ccOp(c.attr[0])
	return func(w ioWriter, s string) (err error) {
		if _, err = w.Write([]byte(cce.set)); nil != err {
			return
		}
		if _, err = w.Write([]byte(s)); nil != err { // write <=> fputs(w, a...)
			return
		}
		_, err = w.Write([]byte(cce.unset))
		return
	}
}
