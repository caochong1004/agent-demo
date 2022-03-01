package channel

import "hci-agent/pkg/agent/model"

type CRChan struct {
	CommandChan  chan *model.Packet
	ResponseChan chan *model.Packet
}

func NewCRChannel(commandSize, RespSize int) *CRChan {
	return &CRChan{
		CommandChan:  make(chan *model.Packet, commandSize),
		ResponseChan: make(chan *model.Packet, RespSize),
	}
}
