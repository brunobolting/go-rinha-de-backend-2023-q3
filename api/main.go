package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/brunobolting/go-rinha-backend/api/handlers"
	"github.com/brunobolting/go-rinha-backend/api/middlewares"
	"github.com/brunobolting/go-rinha-backend/config/env"
	repository "github.com/brunobolting/go-rinha-backend/infra"
	person "github.com/brunobolting/go-rinha-backend/usecase"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASSWORD, env.DB_DATABASE,
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("Error on connect to DB: %s", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Error on ping to DB: %s", err)
	}

	cache := redis.NewClient(&redis.Options{
		Addr:         env.REDIS,
		Password:     "",
		DB:           0,
		PoolSize:     16,
		MinIdleConns: 16,
	})

	r := repository.NewPersonPostgreSql(db)
	c := repository.NewPersonRedis(cache)

	s := person.NewService(r, c)

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	middlewares.MakeMiddlewares(app)

	handlers.MakePersonHandlers(app, s)

	app.Listen(fmt.Sprintf(":%s", env.PORT))
}
