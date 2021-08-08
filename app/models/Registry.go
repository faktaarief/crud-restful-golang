package models

import "github.com/faktaarief/crud-restful-golang/app/database"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: User{}},
		{Model: Post{}},
	}
}

func AddConstraints() {
	database.Connector.Model(Post{}).AddForeignKey(
		"author_id",
		"users(id)",
		"cascade",
		"cascade",
	)
}
