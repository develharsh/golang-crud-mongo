package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/develharsh/golang-crud-mongo/models"
	utilsTypes "github.com/develharsh/golang-crud-mongo/utils"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := utilsTypes.ResponseOfUserCRU{
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

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := utilsTypes.ResponseOfUserCRU{
		Success: true,
		Message: "Signup Successful",
	}
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	uc.session.DB(os.Getenv("MONGO_DATABASE")).C("users").Insert(u)

	w.WriteHeader(http.StatusOK)
	response.UserData = &u
	json.NewEncoder(w).Encode(response)
	// uj, err = json.Marshal(u)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := utilsTypes.ResponseOfUserD{
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

	err := uc.session.DB(os.Getenv("MONGO_DATABASE")).C("users").RemoveId(oid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Message = "No such User Found"
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response.Success = true
	response.Message = "Successfully deleted User"
	json.NewEncoder(w).Encode(response)

}
