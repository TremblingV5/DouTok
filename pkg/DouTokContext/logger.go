package DouTokContext

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"reflect"
)

var (
	DefaultNamespace = zap.Namespace("DouTokDefault")
)

var DefaultLogger = zap.NewNop()

type loggerKey struct{}
type namespaceKey struct{}
type globalFieldKey struct{}

// Extract return a logger from the given context
func Extract(ctx context.Context) *zap.Logger {
	v := ctx.Value(loggerKey{})
	if v == nil {
		logger := DefaultLogger.With(zap.String("logger", "default"))
		return logger
	}

	logger, ok := v.(*zap.Logger)
	if !ok {
		logger := DefaultLogger.With(zap.String("logger", "default"))
		return logger
	}

	logger = logger.With(zap.String("logger", "ctx"))
	return WithContextFields(ctx, logger)
}

// AddLoggerToContext adds the zap.Logger to a context for extraction later.
func AddLoggerToContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// AddFields will add or override fields to context when zap.Namespace is provided.
// It will nest any following fields until another zap.Namespace is provided.
// The difference with pure zap.Namespaces is that when you define a namespace, any following attached
// logger will be nested, not only when you call logger.With
func AddFields(ctx context.Context, fields ...zap.Field) context.Context {
	namespaces := extractNamespaces(ctx)
	globalFields := extractGlobalfields(ctx)

	var namespace zap.Field
	for _, field := range fields {
		if field.Type == zapcore.NamespaceType {
			namespace = field
			continue
		}

		if !reflect.ValueOf(namespace).IsZero() {
			namespaces[namespace] = append(namespaces[namespace], field)
		} else {
			globalFields = append(globalFields, field)
		}
	}

	ctx = context.WithValue(ctx, namespaceKey{}, namespaces)
	ctx = context.WithValue(ctx, globalFieldKey{}, globalFields)

	return ctx
}

// WithContextFields adds context fields to the zap.Logger
func WithContextFields(ctx context.Context, logger *zap.Logger) *zap.Logger {
	ctx = AddFields(
		ctx, zap.String("RequestId", GetRequestID(ctx)),
	)

	// TODO: Add more public fields here.

	logger = logger.With(extractGlobalfields(ctx)...)
	for namespace, fields := range extractNamespaces(ctx) {
		logger = logger.With(Nest(namespace.Key, fields...))
	}

	return logger
}

func extractNamespaces(ctx context.Context) map[zapcore.Field][]zapcore.Field {
	ctxValue := ctx.Value(namespaceKey{})
	namespaces := make(map[zapcore.Field][]zapcore.Field)
	if ctxValue != nil {
		ctxNamespaces, ok := ctxValue.(map[zapcore.Field][]zapcore.Field)
		if ok {
			for k, v := range ctxNamespaces {
				namespaces[k] = v
			}
		}
	}
	return namespaces
}

func extractGlobalfields(ctx context.Context) []zapcore.Field {
	ctxValue := ctx.Value(globalFieldKey{})
	globalFields := make([]zap.Field, 0)
	if ctxValue != nil {
		value, ok := ctxValue.([]zap.Field)
		if ok {
			globalFields = value
		}
	}
	return globalFields
}
