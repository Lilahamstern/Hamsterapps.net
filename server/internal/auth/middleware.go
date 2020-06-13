package auth

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/hamsterapps.net/server/internal/users"
	"github.com/lilahamstern/hamsterapps.net/server/pkg/jwt"
	"net/http"
)

var UserCtxKey = &contextKey{name: "user"}

type contextKey struct {
	name string
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")

		// Allow none authenticated users
		if tokenStr == "" {
			c.Next()
			return
		}

		parsedUser, err := jwt.ParseToken(tokenStr)
		if err != nil {
			err := fmt.Errorf("invalid token")
			_ = c.AbortWithError(http.StatusForbidden, err)
			return
		}

		user, err := users.GetUserByEmail(parsedUser.Email)
		if err != nil {
			c.Next()
			return
		}

		ctx := context.WithValue(c.Request.Context(), UserCtxKey, &user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()

	}
}

//func ContextToGinContext(ctx context.Context) (*gin.Context, error) {
//	ginContext := ctx.Value(UserCtxKey)
//
//	if ginContext == nil {
//		err := fmt.Errorf("could not retrive gin context")
//		return nil, err
//	}
//
//	gc, ok := ginContext.(*gin.Context)
//
//	if !ok{
//		err := fmt.Errorf("wrong type of context")
//		return nil, err
//	}
//
//	return gc, nil
//}

func ForContext(ctx context.Context) (*users.User, error) {
	raw, _ := ctx.Value(UserCtxKey).(*users.User)
	return raw, nil
}
