package api

import (
	"context"
)

type Runtime struct {
	ctx context.Context
}

func (a *Runtime) Startup(ctx context.Context) {
	a.ctx = ctx
}
