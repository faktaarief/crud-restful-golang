package app

import "github.com/faktaarief/crud-restful-golang/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Post{}},
	}
}

func (server *Server) AddConstraints() {
	server.DB.Debug().Model(&models.Post{}).AddForeignKey(
		"author_id",
		"users(id)",
		"cascade",
		"cascade",
	)
}
