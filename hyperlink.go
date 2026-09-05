// SPDX-FileCopyrightText: © 2026 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Created by Jose Luis Tallon <jltallon@w-a-t.group>

package color

import "strings"

const osc8setStart = "\x1b]8;;"
const osc8end = "\x1b\\" // ESC-ST [string terminator]
const osc8resetStart = "\x1b]8;;"

// Hyperlink returns the escaped string corresponding to the provided URL
// (with display text "text"), suitable for display on conforming terminal emulators
// NOTE: no checks for OSC8 support are done -- assumed to be verified previously.
func Hyperlink(url, text string) string {

	var sb strings.Builder
	sb.Grow(len(url) + 12 + len(text)) // make room for escape codes and text

	dumb := checkDumbTerm()
	if !dumb {
		sb.WriteString(osc8setStart)
		sb.WriteString(url)
		sb.WriteString(osc8end)
	}

	////////////////////////////////////////////////
	sb.WriteString(text)
	////////////////////////////////////////////////

	if !dumb {
		sb.WriteString(osc8resetStart)
		sb.WriteString(osc8end)
	}
	return sb.String()
}
