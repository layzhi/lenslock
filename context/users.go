package context

import (
	"context"
	"lenslocked/models"
)

type key string

const (
	userKey key = "user"
)

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *models.User {
	val := ctx.Value(userKey)
	user, ok := val.(*models.User)
	if !ok {
		/* The most likely case is that nothing was ever stored in the context,
		so its doesnt have a type of *models.User. It is also possible that other code in this package wrote
		an invalid value using the user key, so it is important to review code changes in this package
		*/
		return nil
	}
	return user
}
