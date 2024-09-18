from concurrent import futures
import grpc
import recommendation_pb2
import recommendation_pb2_grpc

import component_pb2
import component_pb2_grpc

import pricing_pb2
import pricing_pb2_grpc

import os

class RecommendationServiceServicer(recommendation_pb2_grpc.RecommendationServiceServicer):
    def RecommendBuild(self, request, context):
        budget = request.budget
        selected_component_ids = request.selected_component_ids
        # Implement recommendation logic here
        # For simplicity, return dummy data
        recommended_build = [
            recommendation_pb2.Component(id=1, name="Intel Core i5", category="CPU", price=200.0),
            recommendation_pb2.Component(id=2, name="NVIDIA GTX 1660", category="GPU", price=300.0),
            recommendation_pb2.Component(id=3, name="16GB DDR4 RAM", category="RAM", price=100.0),
        ]
        return recommendation_pb2.RecommendBuildResponse(components=recommended_build)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    recommendation_pb2_grpc.add_RecommendationServiceServicer_to_server(
        RecommendationServiceServicer(), server)
    server.add_insecure_port('[::]:50052')
    print("Recommendation Service gRPC server running on port 50052")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
