package helper

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)

func CreateToken(accountId primitive.ObjectID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["account_id"] = accountId
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	tokenUnsigned := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenSigned, err := tokenUnsigned.SignedString([]byte("KEY"))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return tokenSigned, nil
}
func ExtractTokenFromRequest(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearerToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (string, error) {
	tokenString := ExtractTokenFromRequest(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("KEY"), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["account_id"].(string), nil
	}
	return "", err
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		type Response struct {
			Message string `json:"message"`
			IsValid bool `json:"is_valid"`
		}
		accountId, err := VerifyToken(request)
		if err != nil {
			log.Println("Token is valid")
			json.NewEncoder(writer).Encode(&Response{
				Message: "Invalid token",
				IsValid: false,
			})
			return
		}
		q := request.URL.Query()
		q.Add("account_id",accountId)
		request.URL.RawQuery = q.Encode()
		next.ServeHTTP(writer, request)
	})
}

func GetAccountIDLogged(request *http.Request) primitive.ObjectID {
	accountId, _ := primitive.ObjectIDFromHex(request.URL.Query().Get("account_id"))
	return accountId
}