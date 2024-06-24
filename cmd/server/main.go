package main

import (
	"fmt"

	"github.com/LeVanHieu0509/backend-go/internal/routers"
)

func main() {
	fmt.Println("Startin")

	r := routers.NewsRouter()

	r.Run(":8001")
}
