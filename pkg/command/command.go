package command

import (
	"hci-agent/pkg/agent/model"
	"hci-agent/pkg/util/command"
)

type Func func(w *command.Opts, cmd *model.Packet) ([]*model.Packet, *model.Packet)

var Funcs = FuncMap{}

type FuncMap map[string]Func

func (fs *FuncMap) Add(key string, f Func) {
	p := *fs
	p[key] = f
}
