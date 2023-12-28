package DouTokContext

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	"testing"
)

func TestExtractLogger(t *testing.T) {
	cases := []struct {
		description  string
		given        context.Context
		expected     []zap.Field
		loggerExists bool
	}{
		{
			description: "empty context",
			given:       context.Background(),
			expected:    []zap.Field{zap.String("logger", "default")},
		},
		{
			description: "invalid logger type",
			given:       context.WithValue(context.Background(), loggerKey{}, "invalid"),
			expected:    []zap.Field{zap.String("logger", "default")},
		},
	}

	for _, item := range cases {
		item := item
		t.Run(item.description, func(t *testing.T) {
			ctx := item.given
			if item.loggerExists {
				coreLogger, observedLogs := observer.New(zap.InfoLevel)
				existingLogger := zap.New(coreLogger)
				ctx = context.WithValue(ctx, loggerKey{}, existingLogger)

				extractLogger := Extract(ctx)
				extractLogger.Info("doing log")
				allLogs := observedLogs.All()
				require.ElementsMatch(t, item.expected, allLogs[0].Context)
			}

			expectedLogger := zap.NewNop().With(item.expected...)
			require.Equal(t, expectedLogger, Extract(ctx))
		})
	}
}

func TestAddLoggerToContext(t *testing.T) {
	given := zap.NewExample()
	ctx := AddLoggerToContext(context.Background(), given)

	value, ok := ctx.Value(loggerKey{}).(*zap.Logger)
	require.True(t, ok)
	require.Equal(t, given, value)
}

func TestAddFields_DoesNotModifyOriginalNamespaces(t *testing.T) {
	namespace := zap.Namespace("test_space")
	originalFields := map[zap.Field][]zap.Field{
		namespace: {
			zap.String("original field", "original value"),
		},
	}

	ctx := context.WithValue(context.Background(), namespaceKey{}, originalFields)
	ctx = AddFields(ctx, namespace, zap.String("new field", "new value"))
	require.Equal(t, map[zap.Field][]zap.Field{
		namespace: {
			zap.String("original field", "original value"),
		},
	}, originalFields)

	require.Equal(t, map[zap.Field][]zap.Field{
		namespace: {
			zap.String("original field", "original value"),
			zap.String("new field", "new value"),
		},
	}, ctx.Value(namespaceKey{}))
}
