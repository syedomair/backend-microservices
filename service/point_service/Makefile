.PHONY : 

include .env
export


run_main:
	go run main.go 

clean_docker:
	clear
	docker container rm -f $$(docker ps -aq)
	docker rmi bmc-point_service


test: 
	go test ./...