package kafka

type Config struct {
	Addresses []string

	UseAuthentication bool
	Username          string
	Password          string
}

type Message struct {
	Type MessageType `json:"type"`
	Data interface{} `json:"data"`
}

type MessageType string

const (
	MessageTypeTradeLog MessageType = "trade_log"
	MessageTypeRevert   MessageType = "revert"
)
