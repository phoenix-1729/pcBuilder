package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/yourusername/pcbuilder/backend/api-gateway/proto"
    componentpb "github.com/yourusername/pcbuilder/backend/component-service/proto"
    pricingpb "github.com/yourusername/pcbuilder/backend/pricing-inventory-service/proto"
)

type server struct {
    pb.UnimplementedAPIGatewayServer
    componentClient componentpb.ComponentServiceClient
    pricingClient   pricingpb.PricingInventoryServiceClient
}

func (s *server) BuildPC(ctx context.Context, req *pb.BuildPCRequest) (*pb.BuildPCResponse, error) {
    // Implement logic to fetch components based on the budget, pricing, and inventory.
    components := []string{"CPU", "Motherboard", "GPU"} // Example

    return &pb.BuildPCResponse{Components: components}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50054")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterAPIGatewayServer(grpcServer, &server{})

    log.Println("API Gateway running on port :50054")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
