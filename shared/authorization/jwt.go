package authorization

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT[T UserModelTypes](user T, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(72 * time.Hour).Unix(),
	}
	claims["data"] = user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

func DecodeJWT[T UserModelTypes](tokenString string, jwtSecret string) (T, error) {
	var zero T // zero value of T to return in case of error

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return zero, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		dataClaim, ok := claims["data"]
		if !ok {
			return zero, errors.New("data claim missing")
		}
		jsonBytes, err := json.Marshal(dataClaim)
		if err != nil {
			return zero, err
		}

		var data T
		err = json.Unmarshal(jsonBytes, &data)
		if err != nil {
			return zero, err
		}

		return data, nil
	}

	return zero, errors.New("invalid token")
}
