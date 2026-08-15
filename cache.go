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
