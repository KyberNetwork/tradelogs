package zxrfqv3

import "errors"

var (
	//parser
	ErrInvalidRfqTrade           = errors.New("invalid zeroxv3 rfq")
	ErrNotFoundLogWithEmptyTopic = errors.New("not found log with empty topic")
	ErrDetectRfqButCanNotParse   = errors.New("detect zeroxv3 rfq trade but can not parse")

	//abi
	ErrContractAddressIsZeroAddress = errors.New("contract address is zero address")
)
