package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sekky0905/modern-chat/infra/datastore"
)

func main() {
	db, err := datastore.NewDB()
	if err != nil {
		s := fmt.Sprintf("new database error. err = %v", err)
		panic(s)
	}

	defer func() {
		if err := db.Close(); err != nil {
			// TODO log は後に変更すること
			log.Printf("close database error. err = %v\n", err)
		}
	}()
}
