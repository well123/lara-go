package setup

import (
	"context"
	"github.com/google/gops/agent"
	"goApi/internal/app/config"
	"goApi/pkg/logger"
)

func Monitor(ctx context.Context) func() {
	monitor := config.C.Monitor
	if monitor.Enable {
		err := agent.Listen(agent.Options{
			Addr:            monitor.Addr,
			ConfigDir:       monitor.ConfigDir,
			ShutdownCleanup: false,
		})
		if err != nil {
			logger.WithContext(ctx).Errorf("Agent monitor error: %s", err.Error())
		}
		return func() {
			agent.Close()
		}
	}
	return func() {}
}
