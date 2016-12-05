package context

import (
	"golang.org/x/net/context"

	"github.com/uber-go/zap"

	"os"
	"path"
)

type correlationIdType int

const (
	requestIdKey correlationIdType = iota
	sessionIdKey
)

var logger zap.Logger

// set context.WithValue(ctx, requestIdKey, requestId)
// get ctxRqId, ok := ctx.Value(requestIdKey).(string); ok

// use grpc metadata manage context
// ctx := metadata.NewContext(
//            context.Background(),
//            metadata.Pairs("key1", "val1", "key2", "val2"),
//        )
// md, _ := metadata.FromContext(ctx)
// fmt.Println(md["key1"])

func init() {
	// a fallback/root logger for events without context
	logger = zap.New(
		zap.NewJSONEncoder(zap.TimeFormatter(zap.EpochFormatter("today"))),
		zap.Fields(zap.Int("pid", os.Getpid()),
			zap.String("exe", path.Base(os.Args[0]))),
	)
}

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestIdKey, requestId)
}

// WithSessionId returns a context which knows its session ID
func WithSessionId(ctx context.Context, sessionId string) context.Context {
	return context.WithValue(ctx, sessionIdKey, sessionId)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) zap.Logger {
	newLogger := logger
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(requestIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("rqId", ctxRqId))
		}
		if ctxSessionId, ok := ctx.Value(sessionIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("sessionId", ctxSessionId))
		}
	}
	return newLogger
}
