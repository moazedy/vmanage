package main

import (
	"fmt"
	"vmanage/internal/module/application"
	"vmanage/internal/module/presentation/rest"
	"vmanage/internal/module/repository/pg"
	"vmanage/pkg/infra/config"
	"vmanage/pkg/infra/tx"
	"vmanage/pkg/module/vmanage/application/appservice"
	"vmanage/pkg/module/vmanage/model/entity"
	presentation "vmanage/pkg/module/vmanage/persentation"
	"vmanage/pkg/module/vmanage/repository"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type serverItems struct {
	// basics
	dbConnection *gorm.DB
	txFactory    tx.TXFactory

	// repo layer
	vehicleRepoFactory repository.GenericRepoFactory[entity.Vehicle]

	// app layer
	vehicleAppService appservice.Vehicle

	// presentation layer
	vehiclePresentation presentation.Vehicle

	oauthPresentation presentation.OAuth
}

func initServerItems(cfg config.Config) serverItems {
	items := serverItems{}

	// infra
	items.dbConnection = initDB(cfg.Postgres)
	items.txFactory = tx.NewTXFactory(items.dbConnection)

	// repo layer
	items.vehicleRepoFactory = pg.NewGenericRepoFactory[entity.Vehicle]()

	// app layer
	items.vehicleAppService = application.NewVehicle(items.txFactory, items.vehicleRepoFactory)

	// presentation layer
	items.vehiclePresentation = rest.NewVehicle(items.vehicleAppService)

	items.oauthPresentation = rest.NewOAuth()
	return items
}

// initiating server requirements and start
func startServer() {
	cfg := config.Init()

	serverItems := initServerItems(cfg)

	run(cfg.Server.Address+":"+cfg.Server.Port, serverItems)
}

// opening connection with database
func initDB(cfg config.PostgresConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.ToString()), &gorm.Config{})
	if err != nil {
		println(err)
		panic("failed to connect with db")
	}

	// Check if the desired database exists
	err = db.Exec(
		fmt.Sprintf("CREATE DATABASE %s",
			cfg.Name,
		)).Error
	if err != nil {
		println(err.Error())
	}

	// Reconnect to the newly created database
	db, err = gorm.Open(postgres.Open(cfg.ToStringWithDbName()), &gorm.Config{})
	if err != nil {
		println(err)
		panic("failed to connect with db")
	}

	err = db.Exec(
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`,
	).Error
	if err != nil {
		panic(err.Error())
	}

	// NOTE : register all entities in here
	err = db.AutoMigrate(
		&entity.Vehicle{},
	)
	if err != nil {
		println(err.Error())
		panic("failed to migrate")
	}

	return db
}
