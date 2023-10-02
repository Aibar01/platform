package auth

import (
	"errors"
	"strings"
	"time"
)

type Consumer struct {
	Id     string
	Name   string
	key    string
	Secret string
}

func (c *Consumer) GenerateJWTPayload(ttl int) map[string]any {
	now := time.Now().UTC()
	exp := now.Add(time.Duration(ttl) * time.Second)

	return map[string]any{
		"iss": c.key,
		"iat": now.Unix(),
		"exp": exp.Unix(),
		"consumer": map[string]string{
			"id":   c.Id,
			"name": c.Name,
		},
	}
}

func NewConsumerFromName(name string) (*Consumer, error) {
	jwtConfig, err := getJWTConfig(name)
	if err != nil {
		return nil, err
	}

	consumer := &Consumer{
		Id:     jwtConfig.Id,
		Name:   name,
		key:    jwtConfig.Key,
		Secret: jwtConfig.Secret,
	}

	return consumer, nil
}

func NewConsumerFromToken(tokenString string) (*Consumer, error) {
	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid token")
	}

	if parts[0] != tokenType {
		return nil, errors.New("invalid token type")
	}

	payload, err := extractPayloadFromJWT(parts[1])
	if err != nil {
		return nil, err
	}

	consumer, err := NewConsumerFromName(payload.Consumer.Name)
	if err != nil {
		return nil, err
	}

	if payload.Iss != consumer.key {
		return nil, errors.New("invalid credentials")
	}

	return consumer, nil
}
