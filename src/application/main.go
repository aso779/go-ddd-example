//go:generate go run github.com/99designs/gqlgen

package main

import (
	"context"
	"github.com/aso779/go-ddd-example/application/config"
	"github.com/aso779/go-ddd-example/application/di"
	"github.com/aso779/go-ddd-example/application/serv"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd-example/infrastructure/migrations"
	"go.uber.org/zap"
	"net/http"
)

func main() {
	ctx := context.Background()
	container := di.BuildContainer()

	var initConfig *config.Config
	var initLog *zap.Logger
	var initServ *serv.APIServer

	err := container.Invoke(func(
		conf *config.Config,
		log *zap.Logger,
		serv *serv.APIServer,
		connSet *connection.ConnSet,
	) {
		initConfig = conf
		initLog = log
		initServ = serv

		_ = migrations.Init(ctx, connSet.WritePool())
		err := migrations.Migrate(ctx, connSet.WritePool())
		if err != nil {
			panic(err)
		}
	})

	if err != nil {
		panic(err)
	}

	initLog.Info("server started", zap.String("port", initConfig.Port))
	initLog.Error("server failed", zap.Error(http.ListenAndServe(":"+initConfig.Port, &initServ.MuxCopy)))
}
