package main

import (
	"github.com/amupxm/golang-backend-training/ch9/cfg"
	"github.com/amupxm/golang-backend-training/ch9/domain"
	pg_repository "github.com/amupxm/golang-backend-training/ch9/repository/pg"
	logger "github.com/amupxm/xmus-logger"
)

func main() {
	log := logger.CreateLogger(
		&logger.LoggerOptions{
			LogLevel: 5,
			Std:      true,
		},
	)
	db := pg_repository.CreateConnection(cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_NAME)

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
	log.InformF("user %s crated \n\n", user)
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

	log.InformF("user updated to :: %s \n\n", upU)

	user, err = db.CreateUser(
		&domain.User{
			UserName: "testuser",
			Email:    "testUser@amupxm.com",
			Password: "!@#$%",
		},
	)
	if err != nil {
		panic(err)
	}
	log.InformF("user :: %s crated \n\n", user)

	users, err := db.UsersGetAll()
	if err != nil {
		panic(err)
	}
	log.InformF("Now we have %d users \n", len(*users))

	err = db.DeleteUser(user.Id)
	if err != nil {
		panic(err)
	}

	log.InformF("user with id %d deleted\n", user.Id)
	users, err = db.UsersGetAll()
	if err != nil {
		panic(err)
	}
	log.InformF("Now we have %d users \n", len(*users))

}
