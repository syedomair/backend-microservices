.PHONY : 

include .env_local
export


run_docker:
	unset LOG_LEVEL DATABASE_URL PORT DB DB_MAX_IDLE DB_MAX_OPEN DB_MAX_LIFE_TIME DB_MAX_IDLE_TIME ZAP_CONF GORM_CONF PPROF_ENABLE
	docker compose --env-file .env_local up       

clean_docker:
	docker container rm -f $$(docker ps -aq)
	docker rmi bmc-db
	docker rmi bmc-user_service
	docker rmi bmc-department_service
	docker rmi bmc-point_service

clean_point:
	docker container rm -f $$(docker ps -aq)
	docker rmi bmc-point_service

clean_user:
	docker container rm -f $$(docker ps -aq)
	docker rmi bmc-user_service


test: 
	go test -v ./...

test_race:
	go test ./... -race


protoc_point_v1:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	proto/v1/point/point.proto