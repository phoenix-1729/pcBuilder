load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

go_proto_library(
    name = "component_proto",
    importpath = "github.com/yourusername/component-service/proto",
    proto = ":component_proto_srcs",
)

proto_library(
    name = "component_proto_srcs",
    srcs = ["proto/component.proto"],
)

go_library(
    name = "component_lib",
    srcs = ["main.go"],
    importpath = "github.com/yourusername/component-service",
    deps = [":component_proto"],
)

go_binary(
    name = "component_service",
    embed = ":component_lib",
)
