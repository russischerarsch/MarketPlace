package main

import (
	"context"
	"fmt"
	"log"
	dbconnection "mini-ozon/db_connection"
	"mini-ozon/handlers"
	"mini-ozon/intern/repositories"
	"mini-ozon/intern/services"
)

func main() {
	context, _ := context.WithCancel(context.Background())
	conn, err := dbconnection.CheckConnection(context)
	if err != nil {
		fmt.Println(err)
	}
	ProdRep := repositories.CreateProdRep(conn)
	ProductServ := services.CreateProductService(ProdRep)
	ProdHandler := handlers.NewProductHandler(ProductServ)

	UserRep := repositories.CreateUserRep(conn)
	UserSer := services.CreateUserService(UserRep)
	UserHandler := handlers.CreateUserHandlers(UserSer)

	OrderRep := repositories.CreateOrderRep(conn)
	OrderSer := services.CreateOrderService(OrderRep)
	OrderHandler := handlers.CreateOrderHandlers(OrderSer)

	r := handlers.SetupRouter(ProdHandler, UserHandler, OrderHandler)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
