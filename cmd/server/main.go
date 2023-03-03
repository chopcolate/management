package main

import (
	"context"
	"github.com/chopcolate/management/cmd/server/user"
)

func main() {
	user.Start(context.Background())
}
