package main

import (
	"context"
	"fmt"

	"github.com/cmd-ctrl-q/golang-arch/session"
)

func main() {

	var ctx context.Context = context.Background()

	// set user id in context
	ctx = session.SetUserID(ctx, 42)

	// get user id from context
	fmt.Println(*session.GetUserID(ctx)) // 42

	// set is admin in context
	ctx = session.SetIsAdmin(ctx, true)

	// get is admin value from context
	fmt.Println(*session.GetIsAdmin(ctx)) // true
}
