package main

import (
	"context"
	"fmt"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, favoriteColorKey, "blue")
	anyValue := ctx.Value(favoriteColorKey)
	stringValue, ok := anyValue.(string)
	if !ok {
		fmt.Println(anyValue, "is not a string")
		return
	}

	fmt.Println(stringValue, "is a string")
}
