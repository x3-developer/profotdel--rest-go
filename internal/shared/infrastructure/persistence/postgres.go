package persistence

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"profotdel-rest/config"
)

type Postgres struct {
	*gorm.DB
}

func NewPostgres(cfg *config.Config) *Postgres {
	dsn := GetDSN(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("failed to connect to database: %v", err)
	}

	return &Postgres{db}
}

func GetDSN(cfg *config.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbSsl,
	)
}
