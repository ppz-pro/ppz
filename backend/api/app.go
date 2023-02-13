package api

import (
	"context"
)

type App struct {
	ctx context.Context
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}
