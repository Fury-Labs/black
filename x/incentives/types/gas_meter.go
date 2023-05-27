// Copyright Tharsis Labs Ltd.(Gridiron)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/fury-labs/blackfury/blob/main/LICENSE)

package types

import (
	"github.com/ethereum/go-ethereum/common"
	blackfurytypes "github.com/fury-labs/blackfury/v13/types"
)

// NewGasMeter returns an instance of GasMeter
func NewGasMeter(
	contract common.Address,
	participant common.Address,
	cumulativeGas uint64,
) GasMeter {
	return GasMeter{
		Contract:      contract.String(),
		Participant:   participant.String(),
		CumulativeGas: cumulativeGas,
	}
}

// Validate performs a stateless validation of a Incentive
func (gm GasMeter) Validate() error {
	if err := blackfurytypes.ValidateAddress(gm.Contract); err != nil {
		return err
	}

	return blackfurytypes.ValidateAddress(gm.Participant)
}
