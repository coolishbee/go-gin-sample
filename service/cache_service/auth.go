package cache_service

import (
	"strings"

	"github.com/jameschun7/go-gin-sample/pkg/e"
)

type Account struct {
	UserID    string
	LoginType string
	Token     string
}

func (a *Account) GetAuthKey() string {
	keys := []string{
		e.CACHE_ACCOUNT,
		"LIST",
	}

	if a.UserID != "" {
		keys = append(keys, a.UserID)
	}
	if a.LoginType != "" {
		keys = append(keys, a.LoginType)
	}
	if a.Token != "" {
		keys = append(keys, a.Token)
	}

	return strings.Join(keys, "_")
}
