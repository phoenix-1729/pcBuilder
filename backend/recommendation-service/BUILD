load("@rules_python//python:defs.bzl", "py_binary", "py_library")

py_library(
    name = "recommendation_lib",
    srcs = ["app.py"],
    deps = [
        "@rules_python//python:pip_install",
    ],
)

py_binary(
    name = "recommendation_service",
    srcs = ["app.py"],
    deps = [":recommendation_lib"],
)
