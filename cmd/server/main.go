package main

import (
	"fmt"

	routers "github.com/LeVanHieu0509/backend-go"
)

func main() {
	fmt.Println("Startin")

	r := routers.NewsRouter()

	r.Run(":8001")
}
