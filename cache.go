// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Fatih Arslan, 2013
// SPDX-FileContributor: Jose Luis Tallon <jltallon@w-a-t.group>

package color

import "sync"

var (
	colorsCache   = make(map[ColorAttrib]*Color, 33)
	colorsCacheMu sync.Mutex // protects colorsCache
)

func getCachedColor(a ColorAttrib) *Color {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	//revive:disable:indent-error-flow  False positive from the linter...
	if c, ok := colorsCache[a]; !ok {
		cp := New(a)
		colorsCache[a] = cp
		return cp
	} else {
		return c
	}
}
