package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"../dbUtils"
	"../models"
	"../utils"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateShopClient(w http.ResponseWriter, r *http.Request) {
	shopClient := models.ShopClient{}
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &shopClient)
	utils.CheckError(err)
	if shopClient.Username == "" || shopClient.Password == "" {
		var resp = map[string]interface{}{"status": false, "message": "No password "}
		json.NewEncoder(w).Encode(resp)
		return
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(shopClient.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		var resp = map[string]interface{}{"status": false, "message": "Password Encryption  failed"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	shopClient.Password = string(pass)
	dbUtils.InsertShopClient(shopClient)
	w.WriteHeader(http.StatusAccepted)
}

func Login(w http.ResponseWriter, r *http.Request) {
	request := models.ShopClient{}
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &request)
	utils.CheckError(err)
	shopClient := *dbUtils.SelectShopClient(request.Username)

	expiresAt := time.Now().Add(time.Minute * 10000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(request.Password), []byte(shopClient.Password))

	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	tk := &models.Token{
		Username: shopClient.Username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = shopClient
	json.NewEncoder(w).Encode(resp)
}

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("x-access-token") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.Exception{Message: "Missing auth token"})
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.Exception{Message: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
