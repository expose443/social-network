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
	dao.NewAuthRepository()
	dao.NewCommentRepository()
	dao.NewPostRepository()
	dao.NewReactionRepository()
	dao.NewUserRepository()
	auth := dao.NewAuthRepository()
	service := service.NewService(&dao)
	service.NewAuthService(auth).Register()

}
