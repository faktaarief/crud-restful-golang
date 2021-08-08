package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/faktaarief/crud-restful-golang/app/models"
	"github.com/faktaarief/crud-restful-golang/app/responses"
)

func CreateUser(res http.ResponseWriter, req *http.Request) {
	// body, err := ioutil.ReadAll(req.Body)

	// if err != nil {
	// 	responses.Failed(res, http.StatusUnprocessableEntity, err)
	// }

	user := models.User{}

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		responses.Failed(res, http.StatusUnprocessableEntity, err)
		return
	}

	json.NewEncoder(res).Encode(user)
}
