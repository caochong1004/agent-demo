package command

import (
	"github.com/golang/glog"

	"hci-agent/pkg/agent/model"
)

func NewResponseError(key string, cmdType string, err error) *model.Packet {
	glog.Error(err)
	return &model.Packet{
		Key:     key,
		Type:    cmdType,
		Payload: err.Error(),
	}
}

func NewResponseErrorWithCommit(key string, commit string, cmdType string, err error) *model.Packet {
	glog.Error(err)
	return &model.Packet{
		Key:     key + ".commit:" + commit,
		Type:    cmdType,
		Payload: err.Error(),
	}
}
