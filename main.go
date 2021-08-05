package main

import (
	"context"
	"fmt"

	"github.com/cmd-ctrl-q/golang-arch/session"
)

func main() {

	ctx := context.Background()

	// set context key and value
	ctx = session.SetUserID(ctx, 98765)

	// get userid value
	v := session.GetUserID(ctx)
	fmt.Println("userID key value:", v)

	// get admin value
	b := session.GetAdmin(ctx)
	fmt.Println("admin key value:", b)
}
