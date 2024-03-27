package middlware

import (
	"context"
	"fmt"
	"net/http"
	"testinhousead/internal/config"
	"testinhousead/internal/model"

	"github.com/golang-jwt/jwt"
)

func Aut(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID, ok := r.Context().Value("reqID").(string)
		if !ok {
			reqID = ""
		}

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("----------", err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tknStr := c.Value

		claims := &model.Claims{}

		cfg := config.TokenCFG()

		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (any, error) {
			return []byte(cfg.SecretKey), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				fmt.Println("----------", err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			fmt.Println("----------", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "reqID", reqID)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
