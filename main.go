package main

import (
	"context"
	"fmt"
	"log"
	dbconnection "mini-ozon/db_connection"
	"mini-ozon/handlers"
	"mini-ozon/intern/models/services"
	"mini-ozon/intern/repositories"
)

func main() {
	context, _ := context.WithCancel(context.Background())
	conn, err := dbconnection.CheckConnection(context)
	if err != nil {
		fmt.Println(err)
	}
	ProdRep := repositories.CreateProdRep(conn)
	ProductServ := services.Create(ProdRep)
	ProdHandler := handlers.NewProductHandler(ProductServ)
	r := handlers.SetupRouter(ProdHandler)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
