package worker

import "github.com/KyberNetwork/tradelogs/internal/evmlistenerclient"

type Worker interface {
	Run([]evmlistenerclient.Message) error
}
