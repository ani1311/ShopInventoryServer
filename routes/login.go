package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"../dbUtils"
	"../models"
	"../serverState"
	"../utils"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Login attempt : ")
	request := models.Shop{}
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &request)
	utils.CheckError(err)
	shop := dbUtils.SelectShopWithName(request.Name)
	if shop == nil {
		fmt.Println("No account found ")
		json.NewEncoder(w).Encode(utils.InvalidRequestResponse)
		return
	}

	expiresAt := time.Now().Add(time.Minute * 10000).Unix()
	errf := bcrypt.CompareHashAndPassword([]byte(shop.Password), []byte(request.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("Wrong Password ")
		json.NewEncoder(w).Encode(utils.WrongPasswordResponse)
		return
	}

	tk := &models.Token{
		Username: request.Name,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte(utils.SecretString))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": 200, "message": "Successful"}
	resp[utils.TokenString] = tokenString //Store the token in the response

	sessionID := uuid.New().String()
	serverState.SessionToShopMap[sessionID] = request.Name
	serverState.ShopToSessionMap[request.Name] = sessionID
	resp[utils.SessionIDString] = sessionID //Store the sessionID in the response

	resp[utils.UserString] = request.Name

	fmt.Println("Successful ", resp)

	json.NewEncoder(w).Encode(resp)
}
