#!/bin/bash
#

if [ ! -d "./node_modules" ]; then
  npm install
else
  echo "Node Modules found; skipping npm install"
fi

echo "Generating templ files..."
templ generate

echo "Compiling handlers..."
mkdir -p build/home build/station build/static
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o ./build/home/bootstrap ./lambda/home/main.go
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o ./build/station/bootstrap ./lambda/station/main.go

echo "Bundling files..."
zip home.zip -j ./build/home/bootstrap
zip station.zip -j ./build/station/bootstrap

echo "Compiling tailwindcss..."
npx tailwindcss -i ./internal/assets/tailwind.css -o ./build/static/style.css
