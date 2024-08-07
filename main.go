package main

import (
    "github.com/Tamiru-Alemnew/task-manager/router"
)

func main() {
    r := router.SetupRouter()
    r.Run(":8080") // Run on port 8080
}
