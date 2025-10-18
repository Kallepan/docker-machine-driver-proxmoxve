package main

import (
	"github.com/deepshore/docker-machine-driver-proxmoxve/internal/logger"
	"github.com/rancher/machine/libmachine/drivers/plugin"
)

func main() {
	logger.Init()
	defer logger.CallBackOnExit()

	plugin.RegisterDriver(NewDriver("default", ""))
}
