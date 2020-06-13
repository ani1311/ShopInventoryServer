package utils

import "fmt"

var SuccessfulResponse map[string]interface{}
var InvalidRequestResponse = map[string]interface{}{"status": 400, "message": "Invalid request"}
var PasswordEncryptionFailResponse = map[string]interface{}{"status": 401, "message": "Password Encryption failed"}
var WrongPasswordResponse = map[string]interface{}{"status": 401, "message": "Wrong Password"}
var MissingAuthTokenResponse = map[string]interface{}{"status": 404, "message": "Missing Auth Token"}

const BarcodeString = "barcode"
const SecretString = "secret"
const TokenString = "token"
const SessionIDString = "sessionID"
const OperationString = "operation"
const UserString = "shop"

const AddOperationString = "add"
const RemoveOperationString = "remove"

func init() {
	SuccessfulResponse = map[string]interface{}{"status": 200, "message": "Successful"}
	InvalidRequestResponse = map[string]interface{}{"status": 400, "message": "Invalid request"}
	PasswordEncryptionFailResponse = map[string]interface{}{"status": 401, "message": "Password Encryption failed"}
	WrongPasswordResponse = map[string]interface{}{"status": 401, "message": "Wrong Password"}
	MissingAuthTokenResponse = map[string]interface{}{"status": 404, "message": "Missing Auth Token"}
}

func CheckError(err error) bool {
	if err != nil {
		fmt.Println(err.Error)
		return true
	}
	return false
}
