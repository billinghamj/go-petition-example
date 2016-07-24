package main

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("private-key.pem")
	check(err)

	privateKey, err := loadKey(data)
	check(err)

	now := time.Now()

	claims := map[string]interface{}{
		"iss": "https://example.com",
		"jti": "this-is-definitely-a-uuid",
		"sub": "some-user-id",
		"aud": "some-client-id",
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(time.Hour).Unix(),
	}

	token, err := generateJwt(jwt.SigningMethodES512, privateKey, claims)
	check(err)

	fmt.Println(token)
}

func loadKey(pemData []byte) (crypto.PrivateKey, error) {
	block, _ := pem.Decode(pemData)

	if block == nil {
		return nil, fmt.Errorf("unable to load key")
	}

	if block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("wrong type of key: %s", block.Type)
	}

	return x509.ParseECPrivateKey(block.Bytes)
}

func generateJwt(alg jwt.SigningMethod, key crypto.PrivateKey, claims map[string]interface{}) (string, error) {
	token := jwt.New(alg)
	token.Claims = claims
	return token.SignedString(key)
}
