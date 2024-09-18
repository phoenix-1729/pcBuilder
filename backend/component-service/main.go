package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	pb "component-service/proto"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedComponentServiceServer
	db *sql.DB
}

type Component struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Price         float64 `json:"price"`
	Compatibility string  `json:"compatibility"`
}

func (s *server) GetComponents(ctx context.Context, req *pb.GetComponentsRequest) (*pb.GetComponentsResponse, error) {
	budget := req.GetBudget()
	rows, err := s.db.Query("SELECT id, name, category, price, compatibility FROM components WHERE price <= $1", budget)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var components []*pb.Component
	for rows.Next() {
		var c Component
		if err := rows.Scan(&c.ID, &c.Name, &c.Category, &c.Price, &c.Compatibility); err != nil {
			return nil, err
		}
		components = append(components, &pb.Component{
			Id:            int32(c.ID),
			Name:          c.Name,
			Category:      c.Category,
			Price:         float32(c.Price),
			Compatibility: c.Compatibility,
		})
	}

	return &pb.GetComponentsResponse{Components: components}, nil
}

func (s *server) CheckCompatibility(ctx context.Context, req *pb.CheckCompatibilityRequest) (*pb.CheckCompatibilityResponse, error) {
	componentIDs := req.GetComponentIds()
	// Implement compatibility logic here
	// For simplicity, assume all combinations are compatible
	return &pb.CheckCompatibilityResponse{
		IsCompatible: true,
		Message:      "All selected components are compatible.",
	}, nil
}

func main() {
	// Connect to PostgreSQL
	connStr := "postgres://user:password@postgres/component_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not ping PostgreSQL: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterComponentServiceServer(s, &server{db: db})

	fmt.Println("Component Service gRPC server running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
