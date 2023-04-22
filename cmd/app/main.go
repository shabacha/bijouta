package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/shabacha/pkg/config"
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/infrastructure/datastore"
	"github.com/shabacha/pkg/infrastructure/router"
	"github.com/shabacha/pkg/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Err --> ", err)
	}

	defer sqlDB.Close()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate database")
	}

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
