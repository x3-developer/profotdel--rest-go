package container

import (
	"profotdel-rest/config"
	"profotdel-rest/pkg/database"
	"sync"
)

type Container struct {
	DB *database.DB
	Wg *sync.WaitGroup
}

func NewContainer(cfg *config.Config, wg *sync.WaitGroup) *Container {
	db := database.NewDB(cfg)

	return &Container{
		DB: db,
		Wg: wg,
	}
}
