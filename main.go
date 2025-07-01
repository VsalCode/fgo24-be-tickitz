package main

import (
	"be-cinevo/utils"
	"fmt"
) 

func main() {
	db, err := utils.DBConnect()

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
	}
	defer db.Close()

}