package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func UserGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Service 1 UserGet")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if nil != err {
		log.WithError(err).Error("cannot retrieve URL ID paramater")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(User{ID: id, Email: fmt.Sprintf("example_%d@example.com", id)})
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
