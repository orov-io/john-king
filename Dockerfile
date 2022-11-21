##
##
## STEP 0 - BASE
##
## Install and download dependencies
##
FROM golang:alpine AS base

# Expose debug port
EXPOSE 4000
# Creating a work directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod ./
COPY go.sum ./

# download Go modules and dependencies
RUN go mod download


##
##
## STEP 1 - DEV MODE
##
## Dev mode expects that you mount the service folder in /app as a volume
## It starts a hot reload service with a debugger attadched.
## Be aware, stop the debugger will not clean the debug breakpoints.
##
##

FROM base AS dev

WORKDIR /app

RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT (sleep ${EXPECTED_BUILD_TIME:=30} && dlv --init <(echo "exit -c") connect localhost:4000) & air -c /app/.air.toml

#ENTRYPOINT [ "air" ]

##
## STEP 2 - BUILDER
##

FROM base AS build-env

# Creating a work directory inside the image
WORKDIR /app

# copy directory files
COPY . .

# compile application
RUN go build -o ./app

##
## STEP 2 - DEPLOY
##
FROM alpine as prod

WORKDIR /app

COPY --from=build-env /app/app ./

#TODO: ADD HEALTHCHECK

ENTRYPOINT [ "./app" ]

# ENTRYPOINT ["/bin/ash", "-c", "sleep 100000000"]