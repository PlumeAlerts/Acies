package main

import (
	"fmt"
	"github.com/PlumeAlerts/Acies/migration"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", "root:@/plume?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	//TODO Move to central repo, will prevent issues in the future
	db.LogMode(true)

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		migration.Migrate20181122(),
	})

	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
