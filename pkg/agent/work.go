package agent

import (
	"fmt"
	"github.com/golang/glog"
	"hci-agent/pkg/agent/channel"
	"hci-agent/pkg/agent/model"
	"hci-agent/pkg/command"
	commandutil "hci-agent/pkg/util/command"
	"hci-agent/pkg/websocket"
	"sync"
)

type workerManager struct {
	chans              *channel.CRChan
	clusterId          int

	appClient          websocket.Client

	wg                 *sync.WaitGroup
	stop               <-chan struct{}
	token              string
	platformCode       string
}

func NewWorkerManager(
	chans *channel.CRChan,
	appClient websocket.Client,
	wg *sync.WaitGroup,
	stop <-chan struct{},
	token string,
	platformCode string) *workerManager {
	return &workerManager{
		chans:              chans,
		appClient:          appClient,
		wg:                 wg,
		stop:               stop,
		token:              token,
		platformCode:       platformCode,
	}
}

func (w *workerManager) Start() {
	w.wg.Add(1)
	go w.runWorker()

}

func (w *workerManager) runWorker() {
	defer w.wg.Done()
	for {
		select {
		case <-w.stop:
			glog.Infof("worker down!")
			return
		case cmd := <-w.chans.CommandChan:
			go func(cmd *model.Packet) {
				if cmd == nil {
					glog.Error("got wrong command")
					return

				}
				//vlog.Successf("get command: %s/%s", cmd.Key, cmd.Type)
				var newCmds []*model.Packet = nil
				var resp *model.Packet = nil

				if processCmdFunc, ok := command.Funcs[cmd.Type]; ok {
					opts := &commandutil.Opts{
						StopCh:            w.stop,
						Wg:                w.wg,
						CrChan:            w.chans,
						PlatformCode:      w.platformCode,
						WsClient:          w.appClient,
						Token:             w.token,
					}
					newCmds, resp = processCmdFunc(opts, cmd)
				} else {
					err := fmt.Errorf("type %s not exist", cmd.Type)
					glog.Info(err.Error())
				}

				if newCmds != nil {
					go func(newCmds []*model.Packet) {
						for i := 0; i < len(newCmds); i++ {
							w.chans.CommandChan <- newCmds[i]
						}
					}(newCmds)
				}
				if resp != nil {
					go func(resp *model.Packet) {
						w.chans.ResponseChan <- resp
					}(resp)
				}
			}(cmd)
		}
	}
}

