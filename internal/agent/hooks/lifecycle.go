package hook

import (
	config "github.com/c3os-io/c3os/pkg/config"
	"github.com/c3os-io/c3os/pkg/utils"
)

type Lifecycle struct{}

func (s Lifecycle) Run(c config.Config) error {
	if c.Install.Reboot {
		utils.Reboot()
	}

	if c.Install.Poweroff {
		utils.PowerOFF()
	}
	return nil
}
