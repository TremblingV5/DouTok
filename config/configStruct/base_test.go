package configStruct

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

type BaseWrap struct {
	Wrap Base `envPrefix:"DOUTOK_"`
}

func TestLoad(t *testing.T) {
	config, err := Load[*BaseWrap](context.Background(), &BaseWrap{})
	require.NoError(t, err)

	require.Equal(t, "unknown", config.Wrap.Name)
}
