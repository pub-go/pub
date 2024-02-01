package util

import (
	"context"

	"code.gopub.tech/logs"
)

func Panic(ctx context.Context, arg any) {
	logs.Default().Log(ctx, 1, logs.LevelError, "panic!!!%+v", arg)
	panic(arg)
}
