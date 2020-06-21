package ctx

import "context"

const (
	RequestId = iota
	UserId
	UserToken
)

type PipeFunc func(ctx context.Context) context.Context

func Pipe(ctx context.Context, fns ...PipeFunc) context.Context {
	for _, fn := range fns {
		ctx = fn(ctx)
	}

	return ctx
}

// SetRequestID returns a new Context carrying RequestID.
func RequestID(requestID uint64) PipeFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, RequestId, requestID)
	}
}

// SetUserID returns a new Context with SetUserID.
func SetUserID(userId uint) PipeFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, UserId, userId)
	}
}

// SetUserToken returns a new Context with user JWT
func SetUserToken(token string) PipeFunc {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, UserToken, token)
	}
}
