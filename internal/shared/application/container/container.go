package container

import (
	"profotdel-rest/config"
	"profotdel-rest/internal/shared/infrastructure/persistence"
	"sync"
)

type Container struct {
	PostgresDB *persistence.Postgres
	Wg         *sync.WaitGroup
}

func NewContainer(cfg *config.Config, wg *sync.WaitGroup) *Container {
	pdb := persistence.NewPostgres(cfg)

	return &Container{
		PostgresDB: pdb,
		Wg:         wg,
	}
}
