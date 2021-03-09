// Hardentools
// Copyright (C) 2017-2020 Security Without Borders
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

// IconBase64 is the base 64 encoded hardentools window
var IconBase64 = "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"

// HardentoolsWindowIconStruct is a struct that implements the
// fyne Resource interface
type HardentoolsWindowIconStruct struct {
	NameInt    string
	ContentInt []byte
}

// Name gets the icons name
func (icon HardentoolsWindowIconStruct) Name() string {
	return icon.NameInt
}

// Content gets the icon
func (icon HardentoolsWindowIconStruct) Content() []byte {
	return icon.ContentInt
}
