#!/bin/bash

mkdir -p build/home build/station
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o ./build/home/bootstrap ./lambda/home/main.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o ./build/station/bootstrap ./lambda/station/main.go
zip home.zip -j ./build/home/bootstrap
zip station.zip -j ./build/station/bootstrap
