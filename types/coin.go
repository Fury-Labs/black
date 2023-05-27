// Copyright Tharsis Labs Ltd.(Gridiron)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/fury-labs/blackfury/blob/main/LICENSE)
package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoGridiron defines the default coin denomination used in Gridiron in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Gridiron.
	AttoGridiron string = "afury"

	// BaseDenomUnit defines the base denomination unit for Gridiron.
	// 1 blackfury = 1x10^{BaseDenomUnit} afury
	BaseDenomUnit = 18

	// DefaultGasPrice is default gas price for evm transactions
	DefaultGasPrice = 20
)

// PowerReduction defines the default power reduction value for staking
var PowerReduction = sdkmath.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(BaseDenomUnit), nil))

// NewGridironCoin is a utility function that returns an "afury" coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewGridironCoin(amount sdkmath.Int) sdk.Coin {
	return sdk.NewCoin(AttoGridiron, amount)
}

// NewGridironDecCoin is a utility function that returns an "afury" decimal coin with the given sdkmath.Int amount.
// The function will panic if the provided amount is negative.
func NewGridironDecCoin(amount sdkmath.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoGridiron, amount)
}

// NewGridironCoinInt64 is a utility function that returns an "afury" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewGridironCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoGridiron, amount)
}
