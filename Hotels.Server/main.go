package main

import (
	"fmt"
	"os"

	"github.com/Hotels/src"
	"github.com/Hotels/src/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file err:", err)
	}
	PORT := os.Getenv("PORT")

	if PORT == "" {
		fmt.Println("debes configurar la variable de entorno PORT")
	}

	uri := os.Getenv("URI")

	if err := db.ConexionDB(uri); err != nil {
		fmt.Println("error:", err)
	}

	defer db.DesconectarDB()

	src.AppFiber(PORT)
}
