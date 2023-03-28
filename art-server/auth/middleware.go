package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/humgal/art-server/dao/users"
	"github.com/humgal/art-server/util"
	"github.com/humgal/art-server/util/jwt"
)

var userCtxKey = &contextKey{"user"}
var ipCtxKey = &contextKey{"ip"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			ip, err := util.GetIP(r)
			if err == nil {
				ipstat := users.LoginStatus{IP: ip, LoginTime: time.Now().Format("2006-01-02 15:04:05")}
				r = r.WithContext(context.WithValue(r.Context(), ipCtxKey, &ipstat))
			}
			// Allow unauthenticated users in
			if header == "" {

				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header

			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user := users.User{Username: username}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}

func IpContext(ctx context.Context) *users.LoginStatus {
	raw, _ := ctx.Value(ipCtxKey).(*users.LoginStatus)
	return raw
}
