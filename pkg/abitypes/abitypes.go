package abitypes

import "github.com/ethereum/go-ethereum/accounts/abi"

var (
	Uint256, _ = abi.NewType("uint256", "", nil)
	Bytes32, _ = abi.NewType("bytes32", "", nil)
	Address, _ = abi.NewType("address", "", nil)
	Bytes, _   = abi.NewType("bytes", "", nil)
)
