package main

import (
	"flag"
	"fmt"
	"log"

	tweetpgstore "github.com/alwisisva/twitter-app/internal/tweet/store/postgresql"
	"github.com/alwisisva/twitter-app/pkg/config"
	"github.com/alwisisva/twitter-app/pkg/helper"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var rollback bool
	flag.BoolVar(&rollback, "rollback", false, "")
	var healthCheck bool
	flag.BoolVar(&healthCheck, "check", false, "")
	var env = flag.String("envfile", "env", "enter env file")
	var version = flag.Int("version", 0, "enter version")

	flag.Parse()

	fmt.Println("env", *env)
	if env == nil || *env == "" || *env == "env" {
		config.LoadEnv(".env")
	} else {
		config.LoadEnv(".env.test")
	}

	dbStr := helper.PsqlConnectionString()
	storeClient, err := tweetpgstore.New(dbStr)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(storeClient.Db.DB, &postgres.Config{})
	if err != nil {
		log.Println("err", err)
	}

	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		log.Println("err", err)
	}

	if *version > 0 {
		if err := m.Force(*version); err != nil {
			log.Fatalf("Rollback version failed : %v \n", err)
			return
		}

		log.Printf("Rollback version to %d Succeed \n", *version)
		return
	}

	if healthCheck {
		version, dirty, err := m.Version()
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Printf("Migration status : \n - Version : %d \n - Dirty state : %t \n\n", version, dirty)
		return
	}

	log.Println("Running migration")
	if rollback {
		if err := m.Down(); err != nil {
			log.Fatal(err)
			return
		}
		log.Println("Rollback Done!!!")
	} else {
		if err := m.Up(); err != nil {
			log.Fatal("err migrate up ", err)
			return
		}
		log.Println("Migration Done!!!")
	}
}
