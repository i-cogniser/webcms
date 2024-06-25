package cct

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateRandomKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func GenKey() {
	keyLength := 32 // Длина ключа в байтах (рекомендуемая длина для HMAC-SHA256)
	key, err := generateRandomKey(keyLength)
	if err != nil {
		fmt.Println("Ошибка при генерации ключа:", err)
		return
	}
	fmt.Println("Сгенерированный ключ:", key)
}
