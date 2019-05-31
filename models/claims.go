package models

import "github.com/dgrijalva/jwt-go"

// Claims Struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
