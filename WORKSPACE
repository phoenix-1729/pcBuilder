workspace(name = "pcbuilder")

# Load Go rules and dependencies
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/v0.38.0/rules_go-v0.38.0.tar.gz"],
    sha256 = "b1c90bd23272b8f1e03ad5a8120c5dcfa30303705fda0532685e21520deee8b8",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.28.0/bazel-gazelle-v0.28.0.tar.gz"],
    sha256 = "b51d2c2f2f43af787223a8182a1b66e3e52431e74ed87b97b2b5ecb94161b199",
)

# Go dependencies
load("@io_bazel_rules_go//go:def.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains()

# Python dependencies
http_archive(
    name = "rules_python",
    urls = ["https://github.com/bazelbuild/rules_python/releases/download/0.21.0/rules_python-0.21.0.tar.gz"],
    sha256 = "d4aab8d5de832a835ef520243b85e38db89063b13b4dc8b4386ce1260e73a00b",
)

# gRPC dependencies
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()
