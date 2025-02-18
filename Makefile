.PHONY : 

include .env
export

run_user:
	go run service/user_service/main.go service/user_service/endpoints.go

build_user-srv-docker:
	docker build --progress=plain --no-cache -f service/user_service/Dockerfile -t backend/user-srv-api:latest  \
	--build-arg logLevelEnvVar=$(LOG_LEVEL) \
       --build-arg databaseURLEnvVar=$(DATABASE_URL) \
       --build-arg portEnvVar=$(PORT) \
       --build-arg dBEnvVar=$(DB) \
       --build-arg dBMaxIdleEnvVar=$(DB_MAX_IDLE) \
       --build-arg dBMaxOpenEnvVar=$(DB_MAX_OPEN)\
       --build-arg dBMaxLifeTimeEnvVar=$(DB_MAX_LIFE_TIME) \
       --build-arg dBMaxIdleTimeEnvVar=$(DB_MAX_IDLE_TIME) \
       --build-arg zapConf=$(ZAP_CONF) \
       --build-arg gormConf=$(GORM_CONF) \
       --build-arg pprofEnable=$(PPROF_ENABLE) .
	docker container run  -e LOG_LEVEL='$(LOG_LEVEL)' -e DATABASE_URL='$(DATABASE_URL)' -e PORT=$(PORT) -e DB=$(DB) -e DB_MAX_IDLE=$(DB_MAX_IDLE) -e DB_MAX_OPEN=$(DB_MAX_OPEN) -e DB_MAX_LIFE_TIME=$(DB_MAX_LIFE_TIME) -e DB_MAX_IDLE_TIME=$(DB_MAX_IDLE_TIME) -e ZAP_CONF='$(ZAP_CONF)' -e GORM_CONF=$(GORM_CONF) -e PPROF_ENABLE='$(PPROF_ENABLE)' -p 8185:8185 --name user-srv-api backend/user-srv-api:latest 

build_rm_user-srv-docker:
	docker rm $$(docker ps -a -q)
	docker build --progress=plain --no-cache -f service/user_service/Dockerfile -t backend/user-srv-api:latest  \
	--build-arg logLevelEnvVar=$(LOG_LEVEL) \
       --build-arg databaseURLEnvVar=$(DATABASE_URL) \
       --build-arg portEnvVar=$(PORT) \
       --build-arg dBEnvVar=$(DB) \
       --build-arg dBMaxIdleEnvVar=$(DB_MAX_IDLE) \
       --build-arg dBMaxOpenEnvVar=$(DB_MAX_OPEN)\
       --build-arg dBMaxLifeTimeEnvVar=$(DB_MAX_LIFE_TIME) \
       --build-arg dBMaxIdleTimeEnvVar=$(DB_MAX_IDLE_TIME) \
       --build-arg zapConf=$(ZAP_CONF) \
       --build-arg gormConf=$(GORM_CONF) \
       --build-arg pprofEnable=$(PPROF_ENABLE) .
	docker container run  -e LOG_LEVEL='$(LOG_LEVEL)' -e DATABASE_URL='$(DATABASE_URL)' -e PORT=$(PORT) -e DB=$(DB) -e DB_MAX_IDLE=$(DB_MAX_IDLE) -e DB_MAX_OPEN=$(DB_MAX_OPEN) -e DB_MAX_LIFE_TIME=$(DB_MAX_LIFE_TIME) -e DB_MAX_IDLE_TIME=$(DB_MAX_IDLE_TIME) -e ZAP_CONF='$(ZAP_CONF)' -e GORM_CONF=$(GORM_CONF) -e PPROF_ENABLE='$(PPROF_ENABLE)' -p 8185:8185 --name user-srv-api backend/user-srv-api:latest 

run_rm_user-srv-docker: 
	docker rm $$(docker ps -a -q)
	docker container run  -e LOG_LEVEL='$(LOG_LEVEL)' -e DATABASE_URL='$(DATABASE_URL)' -e PORT=$(PORT) -e DB=$(DB) -e DB_MAX_IDLE=$(DB_MAX_IDLE) -e DB_MAX_OPEN=$(DB_MAX_OPEN) -e DB_MAX_LIFE_TIME=$(DB_MAX_LIFE_TIME) -e DB_MAX_IDLE_TIME=$(DB_MAX_IDLE_TIME) -e ZAP_CONF='$(ZAP_CONF)' -e GORM_CONF=$(GORM_CONF) -e PPROF_ENABLE='$(PPROF_ENABLE)' -p 8185:8185 --name user-srv-api backend/user-srv-api:latest 

run_user-srv-docker: 
	docker container run  -e LOG_LEVEL='$(LOG_LEVEL)' -e DATABASE_URL='$(DATABASE_URL)' -e PORT=$(PORT) -e DB=$(DB) -e DB_MAX_IDLE=$(DB_MAX_IDLE) -e DB_MAX_OPEN=$(DB_MAX_OPEN) -e DB_MAX_LIFE_TIME=$(DB_MAX_LIFE_TIME) -e DB_MAX_IDLE_TIME=$(DB_MAX_IDLE_TIME) -e ZAP_CONF='$(ZAP_CONF)' -e GORM_CONF=$(GORM_CONF) -e PPROF_ENABLE='$(PPROF_ENABLE)' -p 8185:8185 --name user-srv-api backend/user-srv-api:latest 

