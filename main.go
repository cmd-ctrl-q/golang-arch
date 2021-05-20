package main

import (
	"encoding/base64"
	"fmt"
)

// `curl -u username:password -v domain.com`
func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
}
