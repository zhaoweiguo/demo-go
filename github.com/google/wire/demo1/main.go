//+build wireinject

package wire

import (
	"context"
	"github.com/google/wire"
	tmp2 "github.com/zhaoweiguo/demo-go/github.com/google/wire/tmp"
)

func initializeBazwww(ctx context.Context) (tmp2.Baz, error) {
	wire.Build(tmp2.SuperSet)
	return tmp2.Baz{}, nil
}

