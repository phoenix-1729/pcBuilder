load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "pricing_proto",
    importpath = "github.com/yourusername/pcbuilder/backend/pricing-inventory-service/proto",
    proto = ":pricing_proto_srcs",
)

proto_library(
    name = "pricing_proto_srcs",
    srcs = ["proto/pricing.proto"],
)

go_library(
    name = "pricing_lib",
    srcs = ["main.go"],
    importpath = "github.com/yourusername/pcbuilder/backend/pricing-inventory-service",
    deps = [":pricing_proto"],
)

go_binary(
    name = "pricing_service",
    embed = ":pricing_lib",
)
