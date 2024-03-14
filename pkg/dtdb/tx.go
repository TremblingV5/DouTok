package dtdb

import (
	"context"
	"github.com/TremblingV5/DouTok/pkg/dtx"
	"go.uber.org/zap"
	"runtime/debug"
	"strings"
)

type txKey struct{}

type TX interface {
	Commit() error
	Rollback() error
}

type TXHandle struct {
	getTx func() TX
}

func NewTXHandle(getTx func() TX) *TXHandle {
	return &TXHandle{
		getTx: getTx,
	}
}

func WithTx(ctx context.Context, tx TX) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func Tx(ctx context.Context) TX {
	v := ctx.Value(txKey{})
	if v == nil {
		return nil
	}

	tx, ok := v.(TX)
	if !ok {
		return nil
	}
	return tx
}

func persist(ctx context.Context, err error) {
	tx := Tx(ctx)
	logger := dtx.Extract(ctx)
	caller := strings.Split(string(debug.Stack()), "\n")

	if err == nil {
		if err = tx.Commit(); err == nil {
			logger.Info("success commit transaction")
			return
		}

		logger.Error("failed commit transaction", zap.Error(err))
	}

	logger.Error("rolling back failed transaction", zap.Any("caller", caller), zap.Error(err))

	if rollBackErr := tx.Rollback(); rollBackErr != nil {
		logger.Error("failed rollback transaction", zap.Error(rollBackErr))
	}

	logger.Error("rolled back failed transaction", zap.Any("caller", caller), zap.Error(err))
}

func (t *TXHandle) ReplaceTxPersist(ctx context.Context) (context.Context, func(err error)) {
	return t.withTxPersist(ctx, true)
}

func (t *TXHandle) WithTXPersist(ctx context.Context) (context.Context, func(err error)) {
	return t.withTxPersist(ctx, false)
}

func (t *TXHandle) withTxPersist(ctx context.Context, forceReplace bool) (context.Context, func(err error)) {
	t0 := Tx(ctx)
	if t0 != nil && !forceReplace {
		return ctx, func(err error) {
			logger := dtx.Extract(ctx)
			logger.Debug("commit or rollback will be done by parent", zap.Error(err))
		}
	}

	tx := t.getTx()
	ctx = WithTx(ctx, tx)
	return ctx, func(ctx context.Context) func(err error) {
		return func(err error) {
			persist(ctx, err)
		}
	}(ctx)
}
