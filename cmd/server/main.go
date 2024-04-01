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
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalf("You must to enter number(s) of order")
	}

	args := os.Args[1]

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

	var orders []int64
	strArgs := strings.Split(args, ",")
	for _, arg := range strArgs {
		ord, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			log.Fatalf("failed to get arg: %v", err)
		}
		orders = append(orders, ord)
	}

	err = api.PrintOrderByNumber(ctx, orders)
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
}
