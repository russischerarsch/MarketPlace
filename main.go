package main

import (
	"context"
	"fmt"
	"log"
	dbconnection "mini-ozon/db_connection"
	"mini-ozon/handlers"
	"mini-ozon/intern/auth"
	redisclient "mini-ozon/intern/redisClient"
	"mini-ozon/intern/repositories"
	"mini-ozon/intern/services"
	"os"
)

func main() {
	redis := redisclient.NewRedisClient(os.Getenv("REDIS_ADDR"))
	if err := redis.Ping(); err != nil {
		fmt.Println("failed to connect redis")
		return
	}
	context, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := dbconnection.CheckConnection(context)
	if err != nil {
		fmt.Println(err)
	}
	ProdRep := repositories.CreateProdRep(conn)
	ProductServ := services.CreateProductService(ProdRep, redis.Rdb)
	ProdHandler := handlers.NewProductHandler(ProductServ)

	UserRep := repositories.CreateUserRep(conn)
	UserSer := services.CreateUserService(UserRep)
	UserHandler := handlers.CreateUserHandlers(UserSer)

	orderItemRep := repositories.CreateOrderItemRep(conn)

	OrderRep := repositories.CreateOrderRep(conn)
	OrderSer := services.CreateOrderService(OrderRep, ProdRep, orderItemRep)
	OrderHandler := handlers.CreateOrderHandlers(OrderSer)

	AuthSer := auth.CreateAuthService(UserRep)
	AuthHandler := auth.CreateAuthHandlers(AuthSer)

	StatusUpdRep := repositories.CreateChangeRep(conn)
	StatusUpdServ := services.CreateChangeStatServ(StatusUpdRep)
	StatusUpdHandlers := handlers.NewChangeHandler(StatusUpdServ)

	r := handlers.SetupRouter(*AuthHandler, ProdHandler, UserHandler, OrderHandler, StatusUpdHandlers, redis.Rdb)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
