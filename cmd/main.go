package main

import (
	gin2 "github.com/gin-gonic/gin"
	"github.com/pocikode/simple_bank_go/internal/bootstrap"
	"github.com/pocikode/simple_bank_go/internal/delivery/http/route"
)

func main() {
	app := bootstrap.App()
	gin := gin2.Default()

	route.Setup(app.DB, gin)

	gin.Run(":8085")
}
