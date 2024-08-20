package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/LeVanHieu0509/backend-go/internal/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mdb, err := sql.Open("mysql", "levanhieu:levanhieu1234@tcp(127.0.0.1:3306)/shopdevgo")

	if err != nil {
		panic(err)
	}
	defer mdb.Close()

	//excution
	dao := database.New(mdb)

	//get list
	ctx := context.Background()

	err = dao.CreateShop(ctx, "hieu")
	if err != nil {
		log.Fatal(err)
	}

	shopList, err := dao.GetShops(ctx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(shopList)
}
