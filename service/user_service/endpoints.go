package main

import (
	"github.com/syedomair/backend-example/lib/container"
	"github.com/syedomair/backend-example/lib/router"
	"github.com/syedomair/backend-example/service/user_service/user"
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
