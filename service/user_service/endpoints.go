package main

import (
	"github.com/syedomair/backend-microservices/lib/container"
	"github.com/syedomair/backend-microservices/lib/router"
	"github.com/syedomair/backend-microservices/service/user_service/user"
)

func EndPointConf(c container.Container) []router.EndPoint {

	userController := user.Controller{
		Logger: c.Logger(),
		Repo:   user.NewPostgresRepository(c.Db(), c.Logger()),
	}

	return []router.EndPoint{
		{
			Name:        "GetAllUser",
			Method:      router.Get,
			Pattern:     "/users",
			HandlerFunc: userController.GetAllUsers,
		},
	}
}
