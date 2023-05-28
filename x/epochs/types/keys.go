// Copyright 2022 Black Foundation
// This file is part of the Black Network packages.
//
// Black is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Black packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Black packages. If not, see https://github.com/evmos/evmos/blob/main/LICENSE

package types

const (
	// ModuleName defines the module name
	ModuleName = "epochs"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for epochs
	RouterKey = ModuleName
)

// prefix bytes for the epochs persistent store
const (
	prefixEpoch = iota + 1
)

// KeyPrefixEpoch defines prefix key for storing epochs
var KeyPrefixEpoch = []byte{prefixEpoch}
