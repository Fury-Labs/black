// Copyright Tharsis Labs Ltd.(Gridiron)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/fury-labs/blackfury/blob/main/LICENSE)
package upgrade

// The constants used in the upgrade tests are defined here
const (
	// the defaultChainID used for testing
	defaultChainID = "blackfury_9000-1"

	// LocalVersionTag defines the docker image ImageTag when building locally
	LocalVersionTag = "latest"

	// tharsisRepo is the docker hub repository that contains the Gridiron images pulled during tests
	tharsisRepo = "tharsishq/blackfury"

	// upgradesPath is the relative path from this folder to the app/upgrades folder
	upgradesPath = "../../../app/upgrades"

	// versionSeparator is used to separate versions in the INITIAL_VERSION and TARGET_VERSION
	// environment vars
	versionSeparator = "/"
)
