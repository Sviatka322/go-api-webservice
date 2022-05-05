# General Information
it's a simple [gorilla/mux](https://github.com/gorilla/mux) based web service. His main purpose is checking URLs from an imaginary proxy. The web service returns policy value from DB for the requested URL or HTTP 204 (No Content) if such URL is not in the DB.
# Requirements
Service is running on go 1.13 and MongoDB 5.0.8. All require modules describes in go.mod file.
# Docs
Webservise have a swagger docs generated by [swaggo/swag](https://github.com/swaggo/swag). Swagger docs available at the [localhost:10000/swagger/](localhost:10000/swagger/) (for local deployment).
# Setting up environmnet
For local deployment DB using MongoDB in docker compose.
    `cd configs/db`
    `docker-compose up -d`
Then update your MongoDB wtih a sample test data.
    `go run mongo_init.go`
## Install go 1.13
### Mac os
`brew install go@1.13`
# Usage
### Build and rub project
`go build`
`./api-webservice`
or run without building 
`go run main.go`
# Testing
### Unit testing
You can run unit tests.
`go test` in root project directory
### Manual testing
Do manually query web service.
`curl localhost:10000/urlinfo/1/google.com:80/search?query` where **google.com:80/search?query** is the URL we aet going to check
