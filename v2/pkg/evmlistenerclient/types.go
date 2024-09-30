package evmlistenerclient

import (
	"github.com/KyberNetwork/evmlistener/pkg/types"
)

type Message struct {
	ID string
	types.Message
}
