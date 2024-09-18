load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "gateway_proto",
    importpath = "github.com/yourusername/pcbuilder/backend/api-gateway/proto",
    proto = ":gateway_proto_srcs",
)

proto_library(
    name = "gateway_proto_srcs",
    srcs = ["proto/gateway.proto"],
)

go_library(
    name = "gateway_lib",
    srcs = ["main.go"],
    importpath = "github.com/yourusername/pcbuilder/backend/api-gateway",
    deps = [
        ":gateway_proto",
        "//backend/component-service:component_proto",
        "//backend/pricing-inventory-service:pricing_proto",
    ],
)

go_binary(
    name = "gateway_service",
    embed = ":gateway_lib",
)
