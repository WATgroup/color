// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Fatih Arslan, 2013
// SPDX-FileContributor: Jose Luis Tallon <jltallon@w-a-t.group>

package color

import "sync"

type colorcacheEntry struct {
	color Color
	set   string
	unset string
}

type cacheKey = []ColorAttrib

var (
	colorsCache   = make(map[ColorAttrib]colorcacheEntry, 33)
	colorsCacheMu sync.Mutex // protects colorsCache
)

func ccOp(a ColorAttrib) *colorcacheEntry {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	//revive:disable:indent-error-flow  False positive from the linter...
	if ce, ok := colorsCache[a]; !ok {
		cv := New(a)
		cce := colorcacheEntry{cv, cv.str_set(), cv.str_unset()}
		colorsCache[a] = cce
		return &cce
	} else {
		return &ce
	}
}

////////////////////////////////////////////////////////////////////////////////
/// simplified "text string" functions

//revive:disable:var-naming	We like our internal funcs with underscores...

// builds the "set" string for the color
func (c *Color) str_set() string {
	var buf [32]byte
	buf[0] = escRune
	buf[1] = seqbegRune
	ret := buf[:2]
	for i, v := range c.attr {
		if i > 0 {
			ret = append(ret, ';')
		}
		ret = append(ret, smallNum(uint16(v))...)
	}
	ret = append(ret, seqendRune)
	return string(ret)
}

// builds the "unset" string for the color (or plain reset otherwise)
func (c *Color) str_unset() string {
	rr := smallNum(uint16(Reset))
	var buf [40]byte
	buf[0] = escRune
	buf[1] = seqbegRune
	ret := buf[:2]
	for i, v := range c.attr {
		if i > 0 {
			ret = append(ret, ';')
		}
		if ra, ok := mapResetAttributes[v]; ok {
			ret = append(ret, smallNum(uint16(ra))...)
		} else {
			ret = append(ret, rr...)
		}
	}
	ret = append(ret, seqendRune)
	return string(ret)
}
