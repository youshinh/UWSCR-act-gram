package main

import (
	"github.com/zalando/go-keyring"
)

const ServiceName = "UWSCR-act-gram"

// SaveAPIKey は指定されたプロバイダーのAPIキーをセキュアに保存します。
func SaveAPIKey(provider string, key string) error {
	return keyring.Set(ServiceName, provider, key)
}

// GetAPIKey は指定されたプロバイダーのAPIキーをセキュアに読み出します。
func GetAPIKey(provider string) (string, error) {
	key, err := keyring.Get(ServiceName, provider)
	if err != nil {
		if err == keyring.ErrNotFound {
			return "", nil // 未登録の場合は空文字を返す
		}
		return "", err
	}
	return key, nil
}

// DeleteAPIKey は指定されたプロバイダーのAPIキーを削除します。
func DeleteAPIKey(provider string) error {
	err := keyring.Delete(ServiceName, provider)
	if err == keyring.ErrNotFound {
		return nil
	}
	return err
}
