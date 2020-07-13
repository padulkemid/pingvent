package auth

import (
	"context"
	"net/http"
	"strings"

	controller "github.com/padulkemid/pingpos/controllers"
	utils "github.com/padulkemid/pingpos/utils"
)

var userCtxKey = &contextKey{"seller", "username"}
type contextKey struct {
  Role string
  Username string
}


func Middleware() func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(
      func(w http.ResponseWriter, r *http.Request) {
        reqToken := r.Header.Get("authorization")

        // Allow unauthenticated user first
        if reqToken == "" {
          next.ServeHTTP(w, r)
          return
        }

        // Validate token
        splitToken := strings.Split(reqToken, "Bearer ")
        tokenString := splitToken[1]
        data, err := utils.ParseToken(tokenString)

        if err != nil {
          http.Error(w, "Invalid token", http.StatusForbidden)
          return
        }

        // check user
        user, _ := controller.UsernameAdaGak(data.Username)

        // taro context
        ctx := context.WithValue(r.Context(), userCtxKey, user )

        // context baru
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)

      })
  }
}

// finds the user from context
func ForContext(ctx context.Context) (*controller.LoginData, bool){
  data, ok := ctx.Value(userCtxKey).(*controller.LoginData)

  return data, ok
}
