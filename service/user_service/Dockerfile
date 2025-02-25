###########
# Stage 1: Build the Go application
###########
FROM golang:1.23-alpine AS builder

# Define environment variables
ARG logLevelEnvVar=""
ARG databaseURLEnvVar=""
ARG portEnvVar=""
ARG dBEnvVar=""
ARG dBMaxIdleEnvVar=""
ARG dBMaxOpenEnvVar=""
ARG dBMaxLifeTimeEnvVar=""
ARG dBMaxIdleTimeEnvVar=""
ARG zapConf=""
ARG gormConf=""
ARG pprofEnable=""

ENV LOG_LEVEL=$logLevelEnvVar
ENV DATABASE_URL=$databaseURLEnvVar
ENV PORT=$portEnvVar
ENV DB=$dBEnvVar
ENV DB_MAX_IDLE=$dBMaxIdleEnvVar
ENV DB_MAX_OPEN=$dBMaxOpenEnvVar
ENV DB_MAX_LIFE_TIME=$dBMaxLifeTimeEnvVar
ENV DB_MAX_IDLE_TIME=$dBMaxIdleTimeEnvVar
ENV ZAP_CONF=$zapConf
ENV GORM_CONF=$gormConf
ENV PPROF_ENABLE=$pprofEnable

# Install dependencies
RUN apk update && apk upgrade && apk add --no-cache ca-certificates git
RUN update-ca-certificates

RUN mkdir -p /src/service/user_service/user
RUN mkdir -p /src/lib/container
RUN mkdir -p /src/lib/request
RUN mkdir -p /src/lib/response
RUN mkdir -p /src/lib/router
RUN mkdir -p /src/models
RUN mkdir -p /src/config

WORKDIR /src/service/user_service/user
COPY /service/user_service/user/. .

WORKDIR /src/lib
COPY /lib/. .

WORKDIR /src/models
COPY /models/. .

WORKDIR /src/config
COPY /config/. .

WORKDIR /src/
COPY go.mod .
COPY go.sum .
COPY /service/user_service/main.go .
COPY /service/user_service/main_test.go .
COPY /service/user_service/endpoints.go .

# Download Go dependencies
RUN go mod download

# Run tests
RUN go test ./...

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o user-srv-api .

###########
# Stage 2
###########
FROM alpine:latest AS runtime

RUN apk update && apk add --no-cache file

# Expose the port
EXPOSE ${PORT}

WORKDIR /service/user_service/config
COPY --from=builder /src/config .

WORKDIR /service/user_service

# Copy CA certificates from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary from the builder stage
COPY --from=builder /src/user-srv-api user-srv-api

# Debugging: Verify the binary
RUN ls -l && file user-srv-api 

# Set the command to run the application
ENTRYPOINT ["./user-srv-api"]