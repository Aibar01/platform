package auth

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"slices"
	"strings"
)

type UserInterface interface {
	HasPermission(permission string) bool
	hasPermissions(permissions []string) bool
	IsAuthenticated() bool
	IsAdmin() bool
}

type User struct {
	Id          string         `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	PhoneNumber string         `json:"phone_number"`
	Permissions []string       `json:"permissions"`
	Data        map[string]any `json:"data"`
}

func (u *User) HasPermission(permission string) bool {
	return slices.Contains(u.Permissions, permission)
}

func (u *User) hasPermissions(permissions []string) bool {
	for _, permission := range permissions {
		if !slices.Contains(u.Permissions, permission) {
			return false
		}
	}

	return true
}

func (u *User) IsAuthenticated() bool {
	return true
}

func (u *User) IsAdmin() bool {
	if value, ok := u.Data["is_admin"]; ok {
		return value.(bool)
	}

	return false
}

func (u *User) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"id":           u.Id,
		"username":     u.Username,
		"email":        u.Email,
		"phone_number": u.PhoneNumber,
		"permissions":  u.Permissions,
	}

	for key, value := range u.Data {
		data[key] = value
	}

	return json.Marshal(data)
}

func (u *User) UnmarshalJSON(data []byte) error {
	unmarshalledData := make(map[string]any)

	if err := json.Unmarshal(data, &unmarshalledData); err != nil {
		return err
	}

	attrs := []string{"id", "username", "email", "phone_number", "permissions"}

	for _, attr := range attrs {
		if value, ok := unmarshalledData[attr]; ok {
			switch attr {
			case "permissions":
				if permissions, ok := value.([]interface{}); ok {
					for _, permission := range permissions {
						u.Permissions = append(u.Permissions, permission.(string))
					}
				}
			case "id":
				u.Id = value.(string)
			case "username":
				u.Username = value.(string)
			case "email":
				u.Email = value.(string)
			case "phone_number":
				u.PhoneNumber = value.(string)
			}
			delete(unmarshalledData, attr)
		}
	}
	u.Data = unmarshalledData

	return nil
}

type AnonymousUser struct {
	Id          string         `json:"id"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	PhoneNumber string         `json:"phone_number"`
	Permissions []string       `json:"permissions"`
	Data        map[string]any `json:"data"`
}

func (u *AnonymousUser) HasPermission(permission string) bool {
	return false
}

func (u *AnonymousUser) hasPermissions(permissions []string) bool {
	return false
}

func (u *AnonymousUser) IsAuthenticated() bool {
	return false
}

func (u *AnonymousUser) IsAdmin() bool {
	return false
}

func ExtractUserFromToken(tokenString string) (UserInterface, error) {
	consumer, err := NewConsumerFromToken(tokenString)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(tokenString, " ")

	claims := JWTPayload{}
	token, err := jwt.ParseWithClaims(parts[1], &claims, func(token *jwt.Token) (any, error) {
		return []byte(consumer.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	if claims.User != nil {
		return claims.User, nil
	}

	return &AnonymousUser{}, nil
}
