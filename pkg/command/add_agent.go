package command

import (
	"hci-agent/pkg/agent/model"
	"hci-agent/pkg/command/agent"
)

func init() {
	Funcs.Add(model.InitAgent, agent.InitAgent)

	//Funcs.Add(model.ReSyncAgent, agent.ReSyncAgent)
	//
	//Funcs.Add(model.EnvCreate, agent.AddEnv)
	//Funcs.Add(model.EnvDelete, agent.DeleteEnv)
	//Funcs.Add(model.HelmReleaseUpgrade, agent.UpgradeAgent)
	//Funcs.Add(model.AgentUpgrade, agent.UpgradeAgent)
}
