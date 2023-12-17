package middleware

import (
	"account-server/api"
	"account-server/global"
	"account-server/global/constants"
	"account-server/model"
	"account-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	TokenName          = "Authorization" // token请求头key
	TokenPrefix        = "Bearer "       // token前缀
	REnewTokenDuratiom = 10 * 60 * time.Second
)

func tokenErr(c *gin.Context, reset ...string) {
	var msg = "登录已失效"
	if len(reset) > 0 {
		msg = reset[0]
	}
	api.ErrApiResult(c, msg, http.StatusUnauthorized, nil)
	c.Abort()
}

// JWTAuthMiddleware 基于JWT的认证中间件--验证用户是否登录
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(TokenName)
		// token不存在或格式错误
		if authHeader == "" || !strings.HasPrefix(authHeader, TokenPrefix) {
			tokenErr(c, "未登录")
			return
		}
		token := authHeader[len(TokenPrefix):]
		ok, claims := utils.IsTokenValid(token)
		// token解析错误
		if ok == false {
			tokenErr(c)
			return
		}
		stUserId := strconv.Itoa(int(claims.ID))
		stRedisUserTokenKer := strings.Replace(constants.LOGIN_USER_TOKEN_REDIS_KEY, "{ID}", stUserId, -1)
		stRedisToken, err := global.RedisClient.Get(stRedisUserTokenKer)
		// redis token对比
		if err != nil || token != stRedisToken {
			tokenErr(c)
			return
		}
		expireDuration, err := global.RedisClient.GetExpireDuration(stRedisUserTokenKer)
		// 过期
		if err != nil || expireDuration <= 0 {
			tokenErr(c)
			return
		}
		// 有效
		// 续期
		//if expireDuration.Seconds() <= REnewTokenDuratiom.Seconds() {
		//	token, err := service.GenerateTokenAndSetTokenToRedis(claims.ID, claims.Name)
		//	if err != nil {
		//		tokenErr(c)
		//		return
		//	}
		//	c.Header("token", token)
		//}

		c.Set(constants.LOGIN_USER, model.LoginUser{
			ID:       claims.ID,
			Username: claims.Name,
		})

		c.Next()
	}
}
