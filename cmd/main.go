package main

import (
    "github.com/child6yo/forum-sample"
)

func main() {
    srv := new(forum.Server)
    srv.Run("8000", handlers.InitRoutes())
}