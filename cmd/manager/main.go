package main

import (
	goflag "flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"hci-agent/cmd/manager/options"
	"os"
)

func init() {
	goflag.Set("logtostderr", "true")
}
func main()  {
	command := options.NewAgentCommand()

	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	_ = goflag.CommandLine.Parse([]string{})

	defer glog.Flush()

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
