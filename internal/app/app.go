package app

import (
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

	// log.Info(a.Config)

	// log.Info("test")
	// log.Trace("test")
	// log.Debug("test")
	// log.Error("test")
	// log.Warn("test")
	// log.Fatal("test")

}

func (a *app) Setup() {
	var err error
	// var ctx = context.Background()

	// setup configuration
	a.Config = config.New()
	log.Info("configuration loaded...")

	// setup database
	a.DB, err = roach.Init(a.Config)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}
	log.Info("database connected to: ", a.Config.DBHost)

	// a.PGXPool, err = roach.PreparePoolConnection(ctx, a.Config)
	// if err != nil {
	// 	log.Fatal("cannot connect to database: ", err)
	// }
	// log.Info("database connected to: ", a.Config.DBHost)

	// err = a.PGXPool.Ping(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
