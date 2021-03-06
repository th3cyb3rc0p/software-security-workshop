package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jbelmont/software-security/model"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := model.GetUsers()
	payload, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
