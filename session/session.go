package session

import "context"

type stringKey string
type intKey int

var userID stringKey
var isAdmin intKey

func SetUserID(ctx context.Context, value int) context.Context {
	return context.WithValue(ctx, userID, value)
}

func GetUserID(ctx context.Context) int {
	if v := ctx.Value(userID); v != nil {
		// must assert v is an int
		if i, ok := v.(int); ok {
			return i
		}
	}
	return 0
}

func GetAdmin(ctx context.Context) bool {
	if v := ctx.Value(isAdmin); v != nil {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}
