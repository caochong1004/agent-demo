package command

import (
	"hci-agent/pkg/agent/channel"
	"hci-agent/pkg/websocket"
	"sync"
)

type Opts struct {

	StopCh            <-chan struct{}
	Wg                *sync.WaitGroup
	CrChan            *channel.CRChan

	PlatformCode      string
	WsClient          websocket.Client
	Token             string
}
