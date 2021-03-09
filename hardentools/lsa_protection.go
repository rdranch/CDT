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

import (
	"golang.org/x/sys/windows/registry"
)

// LSA contains the registry keys to be hardened.
// For details regarding LSA please refer to:
// https://docs.microsoft.com/en-us/windows-server/security/credentials-protection-and-management/configuring-additional-lsa-protection
var LSA = &RegistrySingleValueDWORD{
	RootKey:         registry.LOCAL_MACHINE,
	Path:            "SYSTEM\\CurrentControlSet\\Control\\Lsa",
	ValueName:       "RunAsPPL",
	HardenedValue:   1,
	shortName:       "LSA",
	longName:        "LSA Protection",
	hardenByDefault: false,
}
