package main

import (
	"context"
	"log"
	"net"
	"strconv"

	pb "pricing-inventory-service/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPricingServiceServer
}

type ComponentPrice struct {
	ID    int32
	Price float32
	Stock int32
}

// Dummy data
var prices = map[int32]ComponentPrice{
	1: {ID: 1, Price: 299.99, Stock: 10},
	2: {ID: 2, Price: 499.99, Stock: 5},
}

func (s *server) GetPricing(ctx context.Context, req *pb.GetPricingRequest) (*pb.GetPricingResponse, error) {
	componentID := req.GetComponentId()
	priceInfo, exists := prices[componentID]
	if !exists {
		return nil, grpc.Errorf(grpc.Code(grpc.ErrServerStopped), "Component not found")
	}
	return &pb.GetPricingResponse{
		Id:    priceInfo.ID,
		Price: priceInfo.Price,
		Stock: priceInfo.Stock,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPricingServiceServer(s, &server{})

	log.Println("Pricing and Inventory Service gRPC server running on port 50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
