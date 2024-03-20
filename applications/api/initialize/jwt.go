package initialize

import (
	"context"
	"github.com/cloudwego/hertz/pkg/protocol"
	"strconv"
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

type LoginResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     string `json:"user_id"`
	Token      string `json:"token"`
}

func InitJwt() {
	AuthMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(ViperConfig.Viper.GetString("JWT.signingKey")),
		Timeout:    12 * time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(string); ok {
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
			// Long 型数据在返回前端时会失真，需使用string类型
			//userId := strconv.FormatInt(claims[constants.IdentityKey].(int64), 10)
			c.SetCookie("token", token, 24*60*60, "/", "", protocol.CookieSameSiteDefaultMode, false, true)

			userId := claims[constants.IdentityKey].(string)
			c.JSON(consts.StatusOK, LoginResp{
				StatusCode: errno.SuccessCode,
				StatusMsg:  errno.Success.ErrMsg,
				UserId:     userId,
				Token:      token,
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
			return claims[constants.IdentityKey].(string)
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
			userIdString := strconv.FormatInt(userId, 10)
			if err == nil && userId != 0 {
				c.Set("JWT_PAYLOAD", jwt.MapClaims{
					constants.IdentityKey: userIdString,
				})
			}
			return userIdString, err
		},
		TokenLookup: "cookie: token, query: token, form: token, param: token",
		TimeFunc:    time.Now,
	})
}

func SetUserId() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		str := c.Keys[constants.IdentityKey].(string)
		if str != "" {
			userId, err := strconv.ParseInt(str, 10, 64)
			if err != nil {
				userId = 0
			}
			c.Set(constants.IdentityKey, userId)
		}
		c.Next(ctx)
	}
}
