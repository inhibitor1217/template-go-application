package loggerfx

import (
	"github.com/inhibitor1217/template-go-application/internal/logger"
	"go.uber.org/fx"
)

var Option = fx.Option(
	fx.Provide(logger.New),
)
