package app

import (
	"fmt"
	"go-boilerplate-api/internal/config"
	roach "go-boilerplate-api/internal/infrastructure/cockroachdb"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type app struct {
	Config *config.Configurations
	DB     *gorm.DB
}

func New() app {
	return app{}
}

func (a *app) Run() {
	a.Setup()

	log.Info(a.Config)

	log.Info("test")
	log.Trace("test")
	log.Debug("test")
	log.Error("test")
	log.Warn("test")
	log.Fatal("test")

}

func (a *app) Setup() {
	var err error

	// setup configuration
	a.Config = config.New()
	log.Info("configuration loaded...")

	// setup database
	a.DB, err = roach.Init(a.Config)
	if err != nil {
		log.Fatal("can't connect to databse: ", err)
	}
	log.Info(fmt.Sprintf("database connected on %s...", a.Config.DBHost))

	// sqlDB, _ := a.DB.DB()
	// sqlDB.Ping()
	// pgxLogger := pool.NewPGXLogrusLogger(&log.Logger{})

	// pgxLogLevel, err := pool.LogLevelFromEnv(a.Config.LogLevel)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// a.PGXPool, err = pool.NewPGXPool(context.Background(), "", pgxLogger, pgxLogLevel)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err = a.PGXPool.Ping(context.Background()); err != nil {
	// 	log.Fatal(err)
	// }

}
