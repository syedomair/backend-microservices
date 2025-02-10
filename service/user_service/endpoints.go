package main

import (
	"backend/lib/router"
	"backend/service/user_service/user"

	"backend/lib/container"
)

func EndPointConf(c container.Container) []router.EndPoint {
	actionController := user.Controller{
		Logger: c.Logger(),
		Repo:   user.NewPostgresRepository(c.Db(), c.Logger()),
	}

	return []router.EndPoint{
		{
			Name:        "GetAllActionsByTask",
			Method:      router.Get,
			Pattern:     "/actions",
			HandlerFunc: actionController.GetAllActions,
		},
	}
}
