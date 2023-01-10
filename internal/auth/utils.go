package auth

import (
	"context"
	wutil "csdlpt/internal/wUtil"
	"encoding/base64"
)

var (
	mocking_secret_key = "mockingsecretkey"
)

func loadPrePrivKey(ctx context.Context) (
	keyBuff []byte, err error) {

	// mocking key
	secretVal := mocking_secret_key

	keyBuff, err = base64.StdEncoding.DecodeString(
		string(secretVal))
	if err != nil {
		err = wutil.NewError(err)
		return
	}
	return
}
