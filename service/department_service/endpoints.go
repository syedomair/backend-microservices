package main

import (
	"github.com/syedomair/backend-microservices/lib/container"
	"github.com/syedomair/backend-microservices/lib/router"
	"github.com/syedomair/backend-microservices/service/department_service/department"
)

func EndPointConf(c container.Container) []router.EndPoint {

	departmentController := department.Controller{
		Logger: c.Logger(),
		Repo:   department.NewDBRepository(c.Db(), c.Logger()),
	}

	return []router.EndPoint{
		{
			Name:        "GetAllDepartments",
			Method:      router.Get,
			Pattern:     "/departments",
			HandlerFunc: departmentController.GetAllDepartments,
		},
	}
}
