package controllers

import (
	"net/http"
)

func (server *Server) Home(res http.ResponseWriter, req *http.Request) {
	// responses.formatter.success(res, http.StatusOK, "This is a Home")
}
