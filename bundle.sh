#!/bin/bash

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap main.go
zip deployment.zip bootstrap
