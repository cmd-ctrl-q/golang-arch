package session

import "context"

type stringKey string
type intKey int

var userID stringKey
var isAdmin intKey

func SetUserID(ctx context.Context, value int) context.Context {
	return context.WithValue(ctx, userID, value)
}

func GetUserID(ctx context.Context) *int {
	if uid := ctx.Value(userID); uid != nil {
		// must assert v is an int,
		// and if so, then return the int
		if v, ok := uid.(int); ok {
			return &v
		}
	}
	return nil
}

// returns context
func SetIsAdmin(ctx context.Context, value bool) context.Context {
	return context.WithValue(ctx, isAdmin, value)
}

// return bool
func GetIsAdmin(ctx context.Context) *bool {
	if isA := ctx.Value(isAdmin); isA != nil {
		// assert the value v is of type bool
		if v, ok := isA.(bool); ok {
			return &v
		}
	}
	return nil

}
