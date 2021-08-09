package main

import (
	"fmt"

	"github.com/amupxm/golang-backend-training/ch8/8.1/cfg"
	"github.com/amupxm/golang-backend-training/ch8/8.1/domain"
	pg_repository "github.com/amupxm/golang-backend-training/ch8/8.1/repository/pg"
)

func main() {
	db := pg_repository.CreateConnection(cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_NAME)
	// us, err := db.UsersGetAll()
	// fmt.Println(us, err)
	user, err := db.CreateUser(
		&domain.User{
			UserName: "amupxm",
			Email:    "amupxm@amupxm.com",
			Password: "!@#$%",
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("created : ", user)
	upU, err := db.UpdateUser(&domain.User{
		UserName: "newAmupxm",
		Email:    "amupxm@amupxm.com",
		Password: "!@#$%",
		Id:       user.Id,
	},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("created : ", upU)

}
