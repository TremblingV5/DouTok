package DouTokContext

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Nest(key string, fields ...zapcore.Field) zapcore.Field {
	return zapcore.Field{
		Key:       key,
		Type:      zapcore.ObjectMarshalerType,
		Interface: Fields(fields),
	}
}

type Fields []zap.Field

func (fs Fields) MarshallLogObject(enc zapcore.ObjectEncoder) error {
	addFields(enc, []zap.Field(fs))
	return nil
}

func addFields(enc zapcore.ObjectEncoder, fields []zap.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}
