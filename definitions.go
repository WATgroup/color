// SPDX-FileCopyrightText: Copyright (c) 2013 Fatih Arslan
// SPDX-FileCopyrightText: © 2025 W-A-T EU Operations Oü
// SPDX-License-Identifier: MPL-2.0
// SPDX-FileContributor: Jose Luis Tallon <jltallon@w-a-t.group>

package color

// Base attributes
const (
	Reset ColorAttrib = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack ColorAttrib = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite

	// used internally for 256 and 24-bit coloring
	foreground
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack ColorAttrib = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack ColorAttrib = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite

	// used internally for 256 and 24-bit coloring
	background
)

// Background Hi-Intensity text colors
const (
	BgHiBlack ColorAttrib = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// Reset attributes -- inverse of the above
const (
	ResetBold ColorAttrib = iota + 22
	ResetItalic
	ResetUnderline
	ResetBlinking
	_
	ResetReversed
	ResetConcealed
	ResetCrossedOut
)

var mapResetAttributes = map[ColorAttrib]ColorAttrib{
	Bold:         ResetBold,
	Faint:        ResetBold,
	Italic:       ResetItalic,
	Underline:    ResetUnderline,
	BlinkSlow:    ResetBlinking,
	BlinkRapid:   ResetBlinking,
	ReverseVideo: ResetReversed,
	Concealed:    ResetConcealed,
	CrossedOut:   ResetCrossedOut,
}
