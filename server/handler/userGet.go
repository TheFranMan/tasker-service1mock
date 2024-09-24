package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Service 1 UserGet")

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(User{ID: 1})
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
