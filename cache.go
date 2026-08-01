// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Fatih Arslan, 2013
// SPDX-FileContributor: Jose Luis Tallon <jltallon@w-a-t.group>

package color

import "sync"

var (
	colorsCache   = make(map[ColorAttrib]*Color, 17)
	colorsCacheMu sync.Mutex // protects colorsCache
)

func getCachedColor(p ColorAttrib) *Color {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	c, ok := colorsCache[p]
	if !ok {
		cv := New(p)
		colorsCache[p] = &cv
	}
	return c
}
