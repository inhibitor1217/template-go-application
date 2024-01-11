package main

import (
	"fmt"

	"github.com/inhibitor1217/template-go-application/internal/env"
	"github.com/inhibitor1217/template-go-application/internal/envfx"
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
	)
}

func logStart(
	env *env.Env,
) {
	fmt.Printf("Running with stage=%s\n", env.Stage)
}
