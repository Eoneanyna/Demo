package saas

import (
	"context"
	"demo/internal/data"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
)

func Parsehost(b *data.BaseInfos) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			fmt.Println(b)
			//if header, ok := transport.FromServerContext(ctx); ok {
			//	token := GetToken(header)
			//	auths := strings.SplitN(token, " ", 2)
			//	if len(auths) != 2 || !strings.EqualFold(auths[0], "Bearer") {
			//		return nil, NullJwtToken
			//	}
			//	out, err2 := ParseToken(ctx, client, header, token)
			//	if err2 != nil {
			//		return nil, ErrWrongContext
			//	}
			//	if err != nil {
			//		return nil, LoginOutJwtToken
			//	}
			//	if out.Body.Data.Userid == 0 || out.Body.Data.Username == "" {
			//		return nil, LoginOutJwtToken
			//	}
			//	ctx = context.WithValue(ctx, "userId", out.Body.Data.Userid)
			ctx = context.WithValue(ctx, "Tenant", "cqcb")
			_, ok := b.Base["cqcb"]
			if !ok {
				db := data.Newcon()
				cleanup := func(ctx2 context.Context) error {
					log.Info("message", "closing the data resources")
					if err := db.Close(); err != nil {
						return err
					}
					return nil
				}
				kratos.AfterStop(cleanup)
			}

			return handler(ctx, req)
			//}
			//return nil, nil
		}
	}
}
