package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golangcodebase/Presentation/api"
)

const (
	defaultPort       = "8420"
	defaultConfigPath = "../../config.yml"
)

var (
	ctx            context.Context
	cancel         context.CancelFunc
	port, confPath string
	engine         *gin.Engine
)

func init() {
	port, confPath = internal.ParseArgs(defaultPort, defaultConfigPath)
	ctx, cancel = context.WithCancel(context.Background())
	var err error

	internal.InitConfig(confPath)

	if err := middleware.SetupSentry(config.Default.Sentry.DSN); err != nil {
		log.Error(err)
	}

	engine = internal.InitEngine(config.Default.Gin.Mode)
	platform.Init(config.Default.Platform)

	database, err = db.New(config.Default.Postgres.URL, config.Default.Postgres.Log)
	if err != nil {
		log.Fatal(err)
	}

	metrics.Setup(database)

	tokenIndexer = tokenindexer.Init(database)
}

func main() {
	//unitOfWork := Repositories.NewUnitOfWork(Persistence.NewDataBaseContext())
	//
	//var name = &DomainEntities.NameEntity{
	//	Id:   1,
	//	Name: "Ehsan",
	//}
	//
	//var atomic = func(uow DomainInterfaces.IUnitOfWork) error {
	//	add, err := uow.NameRepositories().Add(name)
	//	if err != nil {
	//		return err
	//	}
	//
	//	fmt.Println(add)
	//	return nil
	//}
	//find, err := unitOfWork.NameRepositories().GetByName(name.Name)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(find, "-------------------not found")
	//err = unitOfWork.Do(atomic)
	//if err != nil {
	//	log.Println(err)
	//}
	//byName, err := unitOfWork.NameRepositories().Find()
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(byName, "-------------------11111111111")

	api.SetupPlatformAPI(engine)
	api.SetupSwaggerAPI(engine)
}
