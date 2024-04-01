package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	shopApi "github.com/kirillmc/TZ_Go_junior_online_shop/internal/api/shop"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/config"
	"github.com/kirillmc/TZ_Go_junior_online_shop/internal/config/env"
	shopRep "github.com/kirillmc/TZ_Go_junior_online_shop/internal/repository/shop"
	shopSrv "github.com/kirillmc/TZ_Go_junior_online_shop/internal/service/shop"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("You must to enter number(s) of order")
	}

	ctx := context.Background()

	err := config.Load(".env")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	pgCongig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, pgCongig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	shopRepo := shopRep.NewRepository(pool)
	shopServ := shopSrv.NewService(shopRepo)
	api := shopApi.NewImplementation(shopServ)

	err = api.PrintOrderByNumber(ctx, os.Args[1])
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
