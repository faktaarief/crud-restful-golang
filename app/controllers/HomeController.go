package controllers

import (
	"net/http"

	"github.com/faktaarief/crud-restful-golang/app/responses"
)

func Home(res http.ResponseWriter, req *http.Request) {
	responses.Success(res, http.StatusOK, "This is a Home")
}
