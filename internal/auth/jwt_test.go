package auth

import (
	"context"
	"log"
	"testing"
)

func Test_JWT_Full(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	t.Logf("STEP_1: GenerateJWTLoginSession")
	id, login_access, err := GenerateJWTLoginSession(ctx, "user_name", "full_name", "center", "1234567890")
	if err != nil {
		t.Errorf("GenerateJWTLoginSession: %+v\n", err)
		return
	}
	data, status, err := ValidateLoginJWT(ctx, login_access.Token)
	if err != nil {
		t.Errorf("ValidateLoginJWT: %+v\n", err)
		return
	} else {
		t.Logf("STEP_1: ValidateLoginJWT TOKEN= %+v\n", login_access.Token)
		t.Logf("STEP_1: ValidateLoginJWT ID= %+v\n", id)
		t.Logf("STEP_1: ValidateLoginJWT STATUS= %+v\n", status)
		t.Logf("STEP_1: ValidateLoginJWT DATA= %+v\n", data)
	}
	t.Fatal("data", data)
}

func Test_ValidateLoginJWT(t *testing.T) {
	ctx := context.Background()
	jwt_token := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJEYXRhIjp7InNlc3Npb25JZCI6ImEwYTMzZDkwLWExNTMtNDFlNC05ODU0LTI0MjU1ZmZjMWEyZSIsInVzZXJOYW1lIjoiaHRrbiIsImZ1bGxOYW1lIjoiIiwicm9sZSI6IkdJQU5HVklFTiJ9LCJleHAiOjE2NzM0NDM1MTksImp0aSI6IjU2MDBhNDk1LTZkMjEtNDdmMy04YWY0LTUyYzQ4Nzk2MjlmZC02M2JkNjczZmNjNzEyZDUxMWE2MTVlOWUiLCJpYXQiOjE2NzMzNTcxMTl9.8kERg85xjkNhe0xRE0bWevSXvYC-SY7BYxcTC4M1tbUxKT7301RIDah4U4BGRH_qZDkzu21f910bSIwIvLKMfQ"
	token_data, token_status, err := ValidateLoginJWT(ctx, jwt_token)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("token_data", token_data)
	log.Println("token_status", token_status)
	log.Fatal("OK")
}
