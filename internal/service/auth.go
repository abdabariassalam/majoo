package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/abdabariassalam/majoo/constants"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type TokenClaims struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int64     `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy int64     `json:"updated_by"`
	jwt.StandardClaims
}

func (s service) VerifyToken(reqToken string) (*TokenClaims, error) {
	user := &TokenClaims{}
	_, err := jwt.ParseWithClaims(reqToken, user, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.JwtSecret), nil
	})
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, constants.ErrExpiredToken) {
			return nil, constants.ErrExpiredToken
		}
		return nil, constants.ErrInvalidToken
	}
	return user, nil

}

func (s service) Login(userName, password string) (string, error) {
	user, err := s.repo.FindByNameAndPassword(userName, password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("invalid user_name or password")
		}
		return "", err
	}

	claims := TokenClaims{
		user.ID,
		user.Name,
		user.UserName,
		user.CreatedAt,
		user.CreatedBy,
		user.UpdatedAt,
		user.UpdatedBy,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(s.cfg.JWT.JwtExpiration)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.cfg.JWT.JwtSecret))
}
