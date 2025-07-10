package container

import (
	"context"
	"profotdel-rest/config"
	"profotdel-rest/pkg/database"
	"sync"
)

type Container struct {
	DB  *database.DB
	Ctx context.Context
	Wg  *sync.WaitGroup
}

func NewContainer(cfg *config.Config, ctx context.Context, wg *sync.WaitGroup) *Container {
	db := database.NewDB(cfg)

	return &Container{
		DB:  db,
		Ctx: ctx,
		Wg:  wg,
	}
}
