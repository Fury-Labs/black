// Copyright Tharsis Labs Ltd.(Gridiron)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/fury-labs/blackfury/blob/main/LICENSE)

package types

import (
	errorsmod "cosmossdk.io/errors"
)

// RootCodespace is the codespace for all errors defined in this package
const RootCodespace = "blackfury"

// ErrInvalidChainID returns an error resulting from an invalid chain ID.
var ErrInvalidChainID = errorsmod.Register(RootCodespace, 3, "invalid chain ID")
