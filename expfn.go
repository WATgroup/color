// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>


package color


// FprintFnreturns a new function that prints the passed arguments
// colorized as with color.Fprint().
func (c Color) FprintFn() func(w Writer, a ...any) {
	cce := ccOp(c.attr[0])
	return func(w Writer, a ...any) {
		w.Write([]byte(cce.set))
		fprintFn(w, a...)
		w.Write([]byte(cce.unset))
	}
}
