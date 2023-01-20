package src

import (
	"encoding/json"
	"net/http"

	"github.com/bxcodec/faker"
)

func AssignRoutes() {
	http.HandleFunc("/health_check", check)
	http.HandleFunc("/create", create)
	http.HandleFunc("/fetch", fetch)
}

func check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Health check")
}

func create(w http.ResponseWriter, r *http.Request) {
	user := User{}
	err := faker.FakeData(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	err = DbClient().Model(User{}).Create(user).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusCreated)
}

func fetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := make([]User, 0)
	err := DbClient().Model(User{}).Find(&users).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
