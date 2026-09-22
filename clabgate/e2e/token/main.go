package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func main() {
	envFile := flag.String("env-file", "../backend/.env.test", "dotenv file containing JWT_SECRET_KEY_PRIVATE")
	flag.Parse()

	values, err := godotenv.Read(*envFile)
	if err != nil {
		log.Fatal(err)
	}
	privatePEM := strings.ReplaceAll(values["JWT_SECRET_KEY_PRIVATE"], `\n`, "\n")
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privatePEM))
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": "local-smoke", "sub": "42", "aud": "clabgate-local", "azp": "clabgate-local",
		"iat": now.Unix(), "exp": now.Add(5 * time.Minute).Unix(), "nonce": "local-smoke",
		"username": "smoke-student", "email": "smoke@example.invalid", "name": "Smoke Student",
		"roles": []string{"student"}, "last_launch_id": "local-smoke", "k8s:access_type": "user",
	})
	signed, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signed)
}
