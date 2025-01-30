package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/inocencio5117/go-gobid/internal/api"
	"github.com/inocencio5117/go-gobid/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(
		ctx,
		fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
			os.Getenv("GOBID_DATABASE_USER"),
			os.Getenv("GOBID_DATABASE_PASSWORD"),
			os.Getenv("GOBID_DATABASE_HOST"),
			os.Getenv("GOBID_DATABASE_PORT"),
			os.Getenv("GOBID_DATABASE_NAME"),
		))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.Api{
		Router:      chi.NewMux(),
		UserService: services.NewUserService(pool),
	}

	api.BindRoutes()

	fmt.Println("Server started on :3080")

	if err := http.ListenAndServe("localhost:3080", api.Router); err != nil {
		panic(err)
	}
}
