package main

import (
	"log"

	"github.com/expose443/social-network/internal/db"
	"github.com/expose443/social-network/internal/repository"
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
	dao.NewAuthRepository().Register()
}
