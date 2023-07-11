package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	roles []Role `json:"roles"`
	jwt.StandardClaims
}
