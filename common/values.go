package common

import (
	"context"
	"net/http"
)

type traceIDKey struct{}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}

type userNameKey struct{}

func SetUserName(req *http.Request, userName string) *http.Request {
	ctx := context.WithValue(req.Context(), userNameKey{}, userName)
	return req.WithContext(ctx)
}

func GetUserName(ctx context.Context) string {
	name := ctx.Value(userNameKey{})

	if nameStr, ok := name.(string); ok {
		return nameStr
	}
	return ""
}
