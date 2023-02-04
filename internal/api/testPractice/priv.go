package testpractice

import (
	"context"
	"csdlpt/internal/auth"
	"csdlpt/library/ascii"
	"csdlpt/pkg/token"
	"errors"
	"fmt"
	"net/http"
)

/* */
func validateBearer(ctx context.Context, r *http.Request) (int, string, *auth.DataJWT, error) {
	var (
		excute = func(ctx context.Context, r *http.Request) (int, string, *auth.DataJWT, error) {
			var (
				// parseBearerAuth parses an HTTP Bearer Authentication string.
				// "Bearer QWxhZGRpbjpvcGVuIHNlc2FtZQ==" returns QWxhZGRpbjpvcGVuIHNlc2FtZQ.
				parseBearerAuth = func(auth string) (token string, ok bool) {
					const prefix = "Bearer "
					// Case insensitive prefix match. See Issue 22736.
					if len(auth) < len(prefix) || !ascii.EqualFold(auth[:len(prefix)], prefix) {
						return "", false
					}
					return auth[len(prefix):], true
				}
			)
			headerAuth := r.Header.Get("Authorization")
			if len(headerAuth) <= 0 {
				return http.StatusBadRequest, "", &auth.DataJWT{}, errors.New("authorization is empty")
			}
			bearer_token, ok := parseBearerAuth(headerAuth)
			if !ok {
				return http.StatusBadRequest, "", &auth.DataJWT{}, errors.New("authorization is invalid")
			}

			// get from cache DB
			// _, account_id, ok, err := fsdb.CacheLogin.Get(ctx, bearer_token)
			// if err != nil {
			// 	return http.StatusForbidden, bearer_token, &auth.DataJWT{}, err
			// }
			// if !ok {
			// 	return http.StatusForbidden, bearer_token, &auth.DataJWT{}, errors.New("token no login")
			// }
			jwt_data, status, err := auth.ValidateLoginJWT(ctx, bearer_token)
			if err != nil {
				println("ValidateLoginJWT:", err.Error())
			}
			switch status {
			case token.INPUT_EMPTY:
				return http.StatusForbidden, bearer_token, jwt_data, errors.New("token is empty")
			case token.ACCESS_TOKEN_INVALID:
				return http.StatusForbidden, bearer_token, jwt_data, errors.New("token is invalid")
			case token.ACCESS_TOKEN_EXPIRED:
				return http.StatusForbidden, bearer_token, jwt_data, errors.New("token is expired")
			case token.SUCCEED:
				// auth pass
				return http.StatusOK, bearer_token, jwt_data, nil
			default:
				return http.StatusForbidden, bearer_token, jwt_data, errors.New("validate token exception")
			}
		}
	)

	status, token, data, err := excute(ctx, r)
	if err != nil {
		println("[AUTH] ", r.RequestURI, "| Error:", err.Error())
	}
	println("[AUTH] ", r.RequestURI, "| Status:", status)
	println("[AUTH] ", r.RequestURI, "| Token:", token)
	println("[AUTH] ", r.RequestURI, "| Data:", data.UserName)
	println("[AUTH] ", r.RequestURI, "| Access Rights:", fmt.Sprintf("%+v", data))
	return status, token, data, err
}
