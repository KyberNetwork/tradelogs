package zxrfq_v3

import "errors"

var (
	ErrInvalidRfqTrade           = errors.New("invalid zeroxv3 order")
	ErrNotFoundLogWithEmptyTopic = errors.New("not found log with empty topic")
	ErrDetectRfqButCanNotParse   = errors.New("detect zeroxv3 rfq trade but can not parse")
)
