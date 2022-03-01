package agent

import (
	"hci-agent/pkg/agent/model"
	commandutil "hci-agent/pkg/util/command"
)

func InitAgent(opts *commandutil.Opts, cmd *model.Packet) ([]*model.Packet, *model.Packet) {

	response := []byte("初始化成功")
	return nil, &model.Packet{
		Key:     cmd.Key,
		Type:    model.InitAgent,
		Payload: string(response),
	}
}
