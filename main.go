package main

import (
	"be-cinevo/routers"
	"be-cinevo/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := utils.DBConnect()

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}
	defer db.Close()

	r := gin.Default()
	routers.CombineRouters(r)

	fmt.Println("Server starting on port 8080...")
	r.Run(":8080")
}
