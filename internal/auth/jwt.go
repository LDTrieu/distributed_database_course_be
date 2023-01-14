package auth

import (
	"context"
	"csdlpt/pkg/token"
	"csdlpt/pkg/wlog"
	"time"
)

const (
	dur_time = 24 * time.Hour
)

func GenerateJWTLoginSession(ctx context.Context, user_name, full_name, role, center_name, session_id string) (
	id string, token_access *token.JWTDetails, err error) {
	jwtKey, err := loadPrePrivKey(ctx)
	if err != nil {
		return "", nil, err
	}
	return token.GenerateJWT(jwtKey, dur_time,
		&DataJWT{
			SessionID:  session_id,
			UserName:   user_name,
			FullName:   full_name,
			CenterName: center_name,
			Role:       role, // WARNING: role should not be saved in token (RFC 9068)
		})

}

func ValidateLoginJWT(ctx context.Context, jwt_token string) (
	*DataJWT, token.Status, error) {
	jwtKey, err := loadPrePrivKey(ctx)
	if err != nil {
		return nil, token.EXCEPTION, err
	}
	var temp DataJWT
	status, err := token.ValidateJWT(jwtKey, jwt_token, &temp)
	if err != nil {
		return nil, token.EXCEPTION, err
	}
	wlog.Info(ctx, temp)
	return &temp, status, nil

}
