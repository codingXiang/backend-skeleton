package main

import (
	"github.com/codingXiang/backend-skeleton/model"
	. "github.com/codingXiang/backend-skeleton/module/example/delivery/http"
	"github.com/codingXiang/backend-skeleton/module/example/repository"
	"github.com/codingXiang/backend-skeleton/module/example/service"
	"github.com/codingXiang/configer"
	. "github.com/codingXiang/cxgateway/delivery/http"
	"github.com/codingXiang/go-logger"
	"github.com/codingXiang/go-orm"
	"github.com/codingXiang/gogo-i18n"
)

func init() {
	//初始化 Gateway
	Gateway = NewApiGateway("", nil)
	configer.Config.AddCore("storage", configer.NewConfigerCore("yaml", "storage", "./config", "."))
	if data, err := configer.Config.GetCore("storage").ReadConfig(nil); err == nil {
		//設定 Database 連線
		if setting := data.Get("database"); setting != nil {
			orm.NewOrm(orm.InterfaceToDatabase(setting))
			// 建立 Table Schema (Module)
			logger.Log.Debug("create table")
			{
				orm.DatabaseORM.CheckTable(false, gogo_i18n.GoGoi18nMessage{})
				orm.DatabaseORM.CheckTable(false, model.Department{})
				orm.DatabaseORM.CheckTable(false, model.User{})
			}
		} else {
			logger.Log.Error("database setting is not exist")
			panic("must need to setting database config")

		}
		//設定 Redis 連線
		if setting := data.Get("redis"); setting != nil {
			orm.NewRedisClient(orm.InterfaceToRedis(setting))
		} else {
			logger.Log.Error("redis setting is not exist")
			panic("must need to setting redis config")
		}
	}
}

func main() {
	// 建立 Repository (Module)
	var (
		exampleRepo = repository.NewExampleRepository(orm.DatabaseORM.GetInstance())
	)
	// 建立 Service (Module)
	logger.Log.Debug("Create Service Instance")
	var (
		exampleService = service.NewExampleService(exampleRepo)
	)
	// 建立 Handler (Module)
	logger.Log.Debug("Create Http Handler")
	var (
		_ = NewExampleHandler(Gateway, exampleService)
	)
	Gateway.Run()
}
