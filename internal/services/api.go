package services

import (
	"net/http"

	"github.com/BNPrashanth/auth-server/internal/helpers"
	"github.com/BNPrashanth/auth-server/internal/logger"
)

// HandleIndex Function
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(helpers.IndexPage))
}

// Signin Function
func Signin(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Signin called..")
	HandleSignin(w, r)
}

// Welcome Function
func Welcome(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Welcome called..")
	HandleWelcome(w, r)
}

// Refresh Function
func Refresh(w http.ResponseWriter, r *http.Request) {
	logger.Log.Info("Refresh called..")
}
