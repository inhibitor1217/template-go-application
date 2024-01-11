package main

import (
	"github.com/inhibitor1217/template-go-application/internal/env"
	"github.com/inhibitor1217/template-go-application/internal/envfx"
	"github.com/inhibitor1217/template-go-application/internal/logger"
	"github.com/inhibitor1217/template-go-application/internal/loggerfx"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		internalModule(),

		fx.Invoke(logStart),
	).Run()
}

func internalModule() fx.Option {
	return fx.Module(
		"internal",
		envfx.Option,
		loggerfx.Option,
	)
}

func logStart(
	env *env.Env,
	logger logger.Logger,
) {
	logger.Infow("Running application", "stage", env.Stage)
}
