package middleware

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"kratos-demo/utils"
)

func JWTMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				token := tr.RequestHeader().Get("Authorization")
				if token == "" {
					return nil, errors.New("token is empty")
				}
				token, err := utils.CheckToken(token, "Token")
				if err != nil {
					return nil, err
				}

				resultClaims, err := utils.ParseToken(token, &utils.MyCustomClaims{})
				if err != nil {
					return nil, err
				}

				ctx = metadata.NewServerContext(ctx, metadata.New(map[string][]string{
					"username": {resultClaims.Uid},
				}))

				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}
