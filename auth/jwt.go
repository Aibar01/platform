package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"os"
	"strings"
)

const tokenType = "Bearer"

type jwtResponse struct {
	Data []jwtConfigResponse `json:"data"`
}

type jwtConfigResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"Secret"`
	Key    string `json:"key"`
}

type JWTPayload struct {
	Iss      string           `json:"iss"`
	Consumer *ConsumerPayload `json:"consumer"`
	User     *User            `json:"user"`
	jwt.RegisteredClaims
}

type ConsumerPayload struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func getJWTConfig(consumerName string) (*jwtConfigResponse, error) {
	client := &http.Client{}
	kongAdminUrl := os.Getenv("KONG_ADMIN_URL")
	if kongAdminUrl == "" {
		return nil, errors.New("credentials not allowed")
	}

	url := fmt.Sprintf("%s/consumers/%s/jwt", kongAdminUrl, consumerName)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := jwtResponse{}
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	jwtConfig := result.Data[0]

	return &jwtConfig, nil
}

func extractPayloadFromJWT(tokenString string) (*JWTPayload, error) {
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid JWT format")
	}

	payloadPart := parts[1]
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadPart)
	if err != nil {
		return nil, err
	}

	payload := JWTPayload{}
	if err = json.Unmarshal(payloadBytes, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
