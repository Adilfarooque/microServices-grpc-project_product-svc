package main

import (
	"fmt"
	"log"
	"microServices-grpc-project_product-svc/pkg/config"
	"microServices-grpc-project_product-svc/pkg/db"
	"microServices-grpc-project_product-svc/pkg/pb"
	"microServices-grpc-project_product-svc/pkg/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	h := db.Init(c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listen", err)
	}
	fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
