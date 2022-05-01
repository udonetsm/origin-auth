package auth

import (
	"origin-auth/db"
	"origin-auth/getconf"

	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/udonetsm/help/helper"
	"github.com/udonetsm/help/models"
)

func Authorize(w http.ResponseWriter, r *http.Request) {
	auth := models.Auth{}
	json.NewDecoder(r.Body).Decode(&auth)
	ok, user := db.Authentificate(auth)
	if ok {
		token := CreateToken(auth, user, 300)
		w.Write(models.Encode(models.ResponseAuth{Message: token}))
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(&models.ResponseAuth{Error: "invld"})
}

func CreateToken(auth models.Auth, user models.User, livetime int64) string {
	claims := models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + livetime,
			IssuedAt:  time.Now().Unix(),
		},
		User: user,
	}
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	signedToken, err := unsignedToken.SignedString([]byte(getconf.Server.Secret))
	helper.Errors(err, "signedstring(createtoken)")
	return signedToken
}

func Mdlwr(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request got")
		next.ServeHTTP(w, r)
	})
}
