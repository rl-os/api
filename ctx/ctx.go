package ctx

import (
	"context"
	"github.com/rl-os/api/errors"
)

const (
	RequestId = iota
	UserId
	UserToken
)

// prebuild errors
var (
	emptyReqIdErr     = errors.New("internal_context", 500, "empty requestId")
	emptyUserIdErr    = errors.New("internal_context", 401, "empty userId")
	emptyUserTokenErr = errors.New("internal_context", 401, "empty userToken")
)

type PipeFunc func(ctx context.Context) context.Context

func Pipe(ctx context.Context, fns ...PipeFunc) context.Context {
	for _, fn := range fns {
		ctx = fn(ctx)
	}

	return ctx
}

// SetRequestID returns a new Context carrying RequestID.
func SetRequestID(requestID string) PipeFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, RequestId, requestID)
	}
}

// GetRequestId from request context
func GetRequestId(ctx context.Context) (string, error) {
	if id, ok := ctx.Value(RequestId).(string); ok {
		return id, nil
	}

	return "", emptyReqIdErr
}

// SetUserID returns a new Context with SetUserID.
func SetUserID(userId uint) PipeFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, UserId, userId)
	}
}

// GetUserID from request context
func GetUserID(ctx context.Context) (uint, error) {
	if id, ok := ctx.Value(UserId).(uint); ok {
		return id, nil
	}

	return 0, emptyUserIdErr
}

// SetUserToken returns a new Context with user JWT
func SetUserToken(token string) PipeFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, UserToken, token)
	}
}

// GetUserToken from request context
func GetUserToken(ctx context.Context) (string, error) {
	if id, ok := ctx.Value(UserToken).(string); ok {
		return id, nil
	}

	return "", emptyUserTokenErr
}
