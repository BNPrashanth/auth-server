package services

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BNPrashanth/auth-server/internal/logger"
	"github.com/BNPrashanth/auth-server/mock"
	"github.com/BNPrashanth/auth-server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// Create the JWT key used to create the signature
var jwtKey = []byte(viper.GetString("jwtSecret"))

// HandleSignin Function
func HandleSignin(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		logger.Log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := mock.Users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &models.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logger.Log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.Log.Info(tokenString)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Secure:  true,
	})

	data := models.TokenResponse{
		Value:     tokenString,
		Expiresin: expirationTime.String(),
	}
	json.NewEncoder(w).Encode(data)
}
