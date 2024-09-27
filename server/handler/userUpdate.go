package handler

import (
	"fmt"
	"net/http"
)

func UserUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Service 1 UserUpdate")

	w.WriteHeader(http.StatusOK)
}
