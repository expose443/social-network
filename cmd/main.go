package main

import (
	"log"

	"github.com/expose443/social-network/internal/db"
	"github.com/expose443/social-network/internal/repository"
	"github.com/expose443/social-network/internal/service"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	dao := repository.NewRepository(db)

	userCase := service.NewService(&dao)
	authService := userCase.NewAuthService(&dao)
	authService.Register()

}
