package initialize

import (
	"context"
	"strings"
	"time"

	"github.com/hertz-contrib/jwt"

	"github.com/TremblingV5/DouTok/applications/api/biz/model/api"
	"github.com/TremblingV5/DouTok/applications/api/initialize/rpc"
	"github.com/TremblingV5/DouTok/kitex_gen/user"
	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var AuthMiddleware *jwt.HertzJWTMiddleware

func InitJwt() {
	AuthMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(ViperConfig.Viper.GetString("JWT.signingKey")),
		Timeout:    12 * time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) { //nolint
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg //nolint
			default:
				return e.Error()
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			claims := jwt.ExtractClaims(ctx, c)
			userId := claims[constants.IdentityKey].(int64)
			c.JSON(consts.StatusOK, map[string]interface{}{
				"status_code": errno.SuccessCode,
				"status_msg":  errno.Success.ErrMsg,
				"user_id":     userId,
				"token":       token,
			})
		},
		WithNext: func(ctx context.Context, c *app.RequestContext) bool {
			if strings.Contains(string(c.Request.Path()), "feed") {
				var req api.DouyinFeedRequest
				err := c.BindAndValidate(&req)
				if err == nil && req.Token == "" {
					return true
				}
				return false
			}
			//if strings.Contains(string(c.Request.Path()), "publish") {
			//	tokenInForm := string(c.FormValue("token"))
			//
			//	if tokenInForm != "" {
			//		jwt1.Parse(tokenInForm, func(token *jwt1.Token) (interface{}, error) {
			//			if _, ok := token.Method.(*jwt1.SigningMethodRSA); !ok {
			//				return nil, errors.New("unsupported signing method")
			//			}
			//			return
			//		})
			//		return true
			//	}
			//
			//	return false
			//}
			return false
		},
		IdentityKey: constants.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return claims[constants.IdentityKey].(float64)
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"status_code": errno.AuthorizationFailedErrCode,
				"status_msg":  message,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVar api.DouyinUserLoginRequest
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.Username) == 0 || len(loginVar.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			userId, err := rpc.Login(context.Background(), rpc.UserClient, &user.DouyinUserLoginRequest{Username: loginVar.Username, Password: loginVar.Password})
			if err == nil && userId != 0 {
				c.Set("JWT_PAYLOAD", jwt.MapClaims{
					constants.IdentityKey: userId,
				})
			}
			return userId, err
		},
		TokenLookup: "query: token, form: token, param: token",
		TimeFunc:    time.Now,
	})
}
