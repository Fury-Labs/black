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

package incentives

import (
	"sort"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	"github.com/evmos/evmos/v12/x/incentives/keeper"
	"github.com/evmos/evmos/v12/x/incentives/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	data types.GenesisState,
) {
	err := k.SetParams(ctx, data.Params)
	if err != nil {
		panic(errorsmod.Wrapf(err, "error setting params"))
	}

	// Ensure incentives module account is set on genesis
	if acc := accountKeeper.GetModuleAccount(ctx, types.ModuleName); acc == nil {
		panic("the incentives module account has not been set")
	}

	allocationMeters := make(map[string]sdk.Dec)

	for _, incentive := range data.Incentives {
		// Set Incentives
		k.SetIncentive(ctx, incentive)

		// Build allocation meter map
		for _, al := range incentive.Allocations {
			allocationMeters[al.Denom] = allocationMeters[al.Denom].Add(al.Amount)
		}
	}

	// Set allocation meters
	denoms := make([]string, 0, len(allocationMeters))
	for k := range allocationMeters {
		denoms = append(denoms, k)
	}
	sort.Strings(denoms)

	for _, denom := range denoms {
		am := sdk.DecCoin{
			Denom:  denom,
			Amount: allocationMeters[denom],
		}
		k.SetAllocationMeter(ctx, am)
	}

	// Set gas meters
	for _, gasMeter := range data.GasMeters {
		k.SetGasMeter(ctx, gasMeter)
	}
}

// ExportGenesis export module status
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:     k.GetParams(ctx),
		Incentives: k.GetAllIncentives(ctx),
		GasMeters:  k.GetIncentivesGasMeters(ctx),
	}
}
