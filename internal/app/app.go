package app

import (
	"context"
	"goApi/internal/app/config"
	"goApi/internal/app/setup"
	"goApi/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Option func(*options)

type options struct {
	ConfigFile string
	ModelFile  string
	WWWDir     string
	Version    string
}

func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

func SetModelFile(s string) Option {
	return func(o *options) {
		o.ModelFile = s
	}
}

func SetWWWDir(s string) Option {
	return func(o *options) {
		o.WWWDir = s
	}
}

func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func Init(ctx context.Context, opts ...Option) (func(), error) {

	var o options
	for _, opt := range opts {
		opt(&o)
	}
	config.MustLoad(o.ConfigFile)

	if p := o.ModelFile; p != "" {
		config.C.CasBin.Model = p
	}
	if p := o.WWWDir; p != "" {
		config.C.WWW = p
	}

	config.PrintConfigWithJson()

	logger.WithContext(ctx).Infof("Start server.#run_mode %s,#version %s,#pid %d", config.C.RunMode, o.Version, os.Getpid())

	clearLoggerFunc, err := setup.Logger()
	if err != nil {
		return nil, err
	}

	clearMonitorFunc := setup.Monitor(ctx)

	injector, injectorCleanFunc, err := BuildInjector()

	clearHttpServerFunc, err := setup.HttpServer(ctx, injector.Engine)
	return func() {
		clearMonitorFunc()
		clearLoggerFunc()
		clearHttpServerFunc()
		injectorCleanFunc()
	}, nil
}

func Run(ctx context.Context, opt ...Option) error {
	state := 1

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	clearFunc, err := Init(ctx, opt...)
	if err != nil {
		return err
	}
EXIT:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("Receive signal[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}
	clearFunc()
	logger.WithContext(ctx).Infof("Server exit")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
