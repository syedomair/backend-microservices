.PHONY : 

include .env_local
export


run_docker:
	unset LOG_LEVEL DATABASE_URL PORT DB DB_MAX_IDLE DB_MAX_OPEN DB_MAX_LIFE_TIME DB_MAX_IDLE_TIME ZAP_CONF GORM_CONF PPROF_ENABLE
	docker compose --env-file .env_local up       

clean_docker:
	clear
	docker container rm -f $$(docker ps -aq)
	docker rmi bmc-user_service
	docker rmi bmc-department_service


test: 
	go test -v ./...

test_race:
	go test ./... -race