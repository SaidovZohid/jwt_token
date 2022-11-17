package main

import (
	"fmt"
	"log"

	"github.com/SaidovZohid/jwt_token/api"
	"github.com/SaidovZohid/jwt_token/config"
	"github.com/SaidovZohid/jwt_token/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/SaidovZohid/jwt_token/api/docs"
)

func main(){
	cfg := config.Load(".")

	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	cfg.Postgres.Host,
	cfg.Postgres.Port,
	cfg.Postgres.User,
	cfg.Postgres.Password,
	cfg.Postgres.Database,
	)

	psqConn, err := sqlx.Connect("postgres", psql)

	fmt.Println("Configuration: ", psql)
	fmt.Println("Connected Succesfully!")

	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	strg := storage.NewStorageI(psqConn)
	
	apiServer := api.New(&api.RoutetOptions{
		Cfg: &cfg,
		Storage: strg,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}

	log.Print("Server Stopped!")
}