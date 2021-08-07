package app

import "github.com/faktaarief/crud-restful-golang/app/controllers"

func (server *Server) Routes() {
	server.Router.HandleFunc("/", controllers.Home)
}
