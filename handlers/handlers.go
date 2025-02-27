package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"playground/models"
)

// @Summary CreateUser
// @Tag user
// @ID create-user
// @Accept json
// @Product json
// @Success 200 {string} string ok
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	var user models.User
	json.Unmarshal(body, &user)
	user.Create(&user)
	w.Write([]byte("ok"))
}

func FindUsers(w http.ResponseWriter, r *http.Request) {

}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User

	users, err := user.FindAll()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	res, _ := json.Marshal(users)
	w.Write([]byte(res))

}
