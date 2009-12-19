package main

import (
    "hello"
    "web"
)

func main() { web.Run(hello.Routes, "0.0.0.0:9999") }
