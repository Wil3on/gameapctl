//go:build windows
// +build windows

package daemon

import (
	"context"

	"github.com/Wil3on/nordvik_gameap_gameapctl/pkg/utils"
	"github.com/pkg/errors"
)

func Stop(_ context.Context) error {
	err := utils.ExecCommand("winsw", "stop", defaultDaemonConfigPath)
	if err != nil {
		return errors.WithMessage(err, "failed to execute stop gameap-daemon command")
	}

	return nil
}
