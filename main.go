package main

import (
	"context"
	"fmt"
	"time"

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

	ctx.Done()

	// ***** implement context with timeout

	ctx = context.Background()
	// time out at 1000 ms
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	// cancel() is only used to end operations early
	// based on a time out or deadline.
	defer cancel()

	// sleep for 100ms. if sleep loger than timeout, then ctx.Done()
	// is called and stops after sleep is done.
	time.Sleep(5000 * time.Millisecond)

	select {
	case <-ctx.Done():
		fmt.Println("work not finished")
	default:
		fmt.Println("work done")
	}

	// ***** launch 100 go routines

}
