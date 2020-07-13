package auth

import (
  "context"
  "net/http"

  controller "github.com/padulkemid/pingpos/controllers"
  utils "github.com/padulkemid/pingpos/utils"
)

type contextKey struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

var userCtxKey = &contextKey{
  ID: "id",
  Username: "username",
}

func Middleware() func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(
      func(w http.ResponseWriter, r *http.Request) {
        c, err := r.Cookie("token")

        // Allow unauthenticated user first
        if err != nil || c == nil {
          next.ServeHTTP(w, r)
          return
        }

        // Validate token
        tokenString := c.Value
        data, err := utils.ParseToken(tokenString)

        if err != nil {
          http.Error(w, "Invalid token", http.StatusForbidden)
          return
        }

        // check user
        checkUname, err := controller.UsernameAdaGak(data.Username)
        if err != nil {
          next.ServeHTTP(w, r)
          return
        }

        // taro context
        ctx := context.WithValue(r.Context(), userCtxKey, checkUname )

        // context baru
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)

      })
  }
}

// finds the user from context
func ForContext(ctx context.Context) *controller.LoginData {
  data, _ := ctx.Value(userCtxKey).(*controller.LoginData)

  return data
}
