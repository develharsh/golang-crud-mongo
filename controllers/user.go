package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/develharsh/golang-crud-mongo/models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

type ResponseOfGetUser struct {
	Success  bool         `json:"success"`
	Message  string       `json:"message"`
	UserData *models.User `json:"userData"`
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := ResponseOfGetUser{
		Success: false,
	}
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		response.Message = "No such User Found"
		json.NewEncoder(w).Encode(response)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}
	err := uc.session.DB(os.Getenv("MONGO_DATABASE")).C("users").FindId(oid).One(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Message = "No such User Found"
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response.Success = true
	response.Message = "Successfully fetched user info"
	response.UserData = &u
	json.NewEncoder(w).Encode(response)

}

// CreateUser
// DeleteUser
