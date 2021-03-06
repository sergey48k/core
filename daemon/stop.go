package daemon

import (
	"github.com/mesg-foundation/core/container"
)

// Stop stops the MESG Core docker container.
func Stop() error {
	status, err := Status()
	if err != nil || status == container.STOPPED {
		return err
	}
	return defaultContainer.StopService(Namespace())
}
