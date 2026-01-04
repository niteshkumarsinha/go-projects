package login

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateToken(username string) (string, error) {
	expiration := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expiration),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))	
	if err != nil {
		return "", err
	}
	return signedToken, nil
}	

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}	

	valid := credentials.Username == "admin" && credentials.Password == "admin123"
	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}	

	tokenString, err := GenerateToken(credentials.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}	
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	
}	