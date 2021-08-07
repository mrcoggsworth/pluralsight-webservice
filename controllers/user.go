package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/mrcoggsworth/pluralsight-webservice/models"
)

type userController struct {
	userIdPattern *regexp.Regexp
}

func (uc userController) ParseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users"{
		switch r.Method {
		case http.MethodGet:
			uc.GetAll(w, r)
		case http.MethodPost:
			uc.Post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIdPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			uc.Get(id, w)
		case http.MethodPut:
			uc.Put(id, w, r)
		case http.MethodDelete:
			uc.Delete(id, w)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func (uc userController) GetAll(w http.ResponseWriter, r *http.Request) {
	EncodeResponseAsJson(models.GetUsers(), w)
}

func (uc userController) Get(id int, w http.ResponseWriter) {
	u, err := models.GetUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	EncodeResponseAsJson(u, w)
}

func (uc userController) Post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.ParseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	EncodeResponseAsJson(u, w)

}

func (uc userController) Put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.ParseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}

	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	EncodeResponseAsJson(u, w)
}

func (uc userController) Delete(id int, w http.ResponseWriter) {
	err := models.RemoveUser(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func newUserController() *userController {
	return &userController{
		userIdPattern: regexp.MustCompile(`^/users/(\d+/?)`),
	}
}
