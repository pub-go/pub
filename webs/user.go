package webs

import (
	"crypto/md5"
	"fmt"
	"time"

	"code.gopub.tech/logs"
	"code.gopub.tech/logs/pkg/kv"
	"code.gopub.tech/pub/dal/caches"
	"code.gopub.tech/pub/dal/model"
	"code.gopub.tech/pub/dal/query"
	"code.gopub.tech/pub/dto"
	"code.gopub.tech/pub/util"
	"github.com/gin-gonic/gin"
)

var cache = caches.GetDefaultCache()

const LoginCookieName = "login"

// SetUserInfo 登录成功时将登录信息设置到缓存和 cookie 中
func SetUserInfo(ctx *gin.Context, loginInfo *dto.LoginInfo) {
	key := fmt.Sprintf("%x", md5.Sum(util.RandBytes(16)))
	cache.Set(ctx, key, loginInfo)
	ctx.SetCookie(LoginCookieName, key, int(7*24*time.Hour.Seconds()), "/", "",
		ctx.Request.URL.Scheme == "https", true)
}

// LoginInfo 中间件 从 Cookie 中获取登录信息
func LoginInfo(ctx *gin.Context) {
	if key, err := ctx.Cookie(LoginCookieName); err == nil {
		if val, err := cache.Get(ctx, key); err == nil {
			if info, ok := val.(*dto.LoginInfo); ok {
				ua := ctx.Request.UserAgent()
				if ua != info.UserAgent {
					logs.Warn(ctx, "UserAgent mismatch: got %v, want %v", ua, info.UserAgent)
				} else {
					ctx.Set(LoginCookieName, info)
					ctx.Set("user", GetUser(ctx))
				}
			}
		}
	}
	ctx.Next()
}

const UserKey = "currentUser"

// GetUser 用从 Cookie 解析到的登录信息，去查数据库中用户信息
func GetUser(ctx *gin.Context) *dto.UserInfo {
	if val, ok := ctx.Get(UserKey); ok {
		if result, ok := val.(*dto.UserInfo); ok {
			return result
		}
	}
	val, ok := ctx.Get(LoginCookieName)
	if !ok {
		return nil
	}
	info, ok := val.(*dto.LoginInfo)
	if !ok {
		return nil
	}
	u := query.User
	user, err := u.WithContext(ctx).Where(u.Username.Eq(info.Username)).First()
	if err != nil {
		return nil
	}
	user.Salt = ""
	user.Password = nil

	result := &dto.UserInfo{
		User: user,
	}

	m := query.UserMeta
	if metas, err := m.WithContext(ctx).Where(m.UserID.Eq(user.ID), m.Key.Eq(model.UserMetaKeyRole)).Find(); err == nil {
		for _, meta := range metas {
			result.Roles = append(result.Roles, meta.Value)
		}
	}
	ctx.Set(UserKey, result)
	SetContext(ctx, kv.Add(GetContext(ctx), "user", user.Username))
	return result
}
