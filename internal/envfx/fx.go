package envfx

import (
	"github.com/inhibitor1217/template-go-application/internal/env"
	"go.uber.org/fx"
)

var Option = fx.Options(
	fx.Invoke(env.Init),
	fx.Provide(env.Load),
)
