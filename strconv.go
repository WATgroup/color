// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package color

const digits = "0123456789abcdefghijklmnopqrstuvwxyz"

// const nSmalls = 100

// smalls is the formatting of 00..99 concatenated.
// It is then padded out with 56 x's to 256 bytes,
// so that smalls[x&0xFF] has no bounds check.
const smalls = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"

// small returns the string for an i with 0 <= i < nSmalls.
func smallNum(i uint16) string {
	if i < 10 {
		return digits[i : i+1]
	}
	return smalls[i*2 : i*2+2]
}
