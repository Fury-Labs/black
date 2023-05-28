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

// DefaultGenesisState sets default fee market genesis state.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:   DefaultParams(),
		BlockGas: 0,
	}
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(params Params, blockGas uint64) *GenesisState {
	return &GenesisState{
		Params:   params,
		BlockGas: blockGas,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}
