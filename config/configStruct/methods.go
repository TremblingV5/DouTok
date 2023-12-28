package configStruct

import (
	"context"
	"github.com/caarlos0/env/v9"
)

func Load[T any](ctx context.Context, config T) (T, error) {
	if err := env.Parse(config); err != nil {
		return config, err
	}
	return config, nil
}
