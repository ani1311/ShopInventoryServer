package routes

import (
	"encoding/json"
	"net/http"

	"../utils"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test success"))
	json.NewEncoder(w).Encode(utils.SuccessfulResponse)
}
