package main

import (
	"database/sql"
	"fc3-grpc/internal/database"
	"fc3-grpc/internal/pb"
	"fc3-grpc/internal/service"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection" // Import the reflection package
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer) // Register the reflection service
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
