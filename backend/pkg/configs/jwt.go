package configs

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"
)

type JWTRSAKeyConfig struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Expire     time.Duration
}

type JWTSecretKeyConfig struct {
	Secret string
	Expire time.Duration
}

type JWTConfig struct {
	AccessKey  *JWTRSAKeyConfig
	RefreshKey *JWTRSAKeyConfig
}

func NewJWTKeyConfig(privateKeyStr string, publicKeyStr string, expire time.Duration) (*JWTRSAKeyConfig, error) {
	// Нормализуем строки с ключами (заменяем \n на реальные переносы строк)
	privateKeyStr = normalizePEM(privateKeyStr)
	publicKeyStr = normalizePEM(publicKeyStr)

	privateKey, err := parsePrivateKey(privateKeyStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey, err := parsePublicKey(publicKeyStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	return &JWTRSAKeyConfig{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Expire:     expire,
	}, nil
}

// normalizePEM преобразует строку с \n в корректный PEM-формат
func normalizePEM(pemString string) string {
	return strings.ReplaceAll(pemString, `\n`, "\n")
}

func parsePrivateKey(pemString string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemString))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing private key")
	}

	// Пробуем разные форматы приватных ключей
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err == nil {
		return priv, nil
	}

	// Пробуем PKCS8
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key (tried PKCS1 and PKCS8): %w", err)
	}

	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA private key")
	}

	return rsaKey, nil
}

func parsePublicKey(pemString string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemString))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return rsaPub, nil
}
