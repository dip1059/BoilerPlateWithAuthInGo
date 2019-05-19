package main

import (
	Mig "BoilerPlateWithAuthInGo/Database/Migrations"
	Seed "BoilerPlateWithAuthInGo/Database/Seeders"
	"BoilerPlateWithAuthInGo/Routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	Mig.Migrate()
	Seed.Seed()
	Routes.Routes()
}
