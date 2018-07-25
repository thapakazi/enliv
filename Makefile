-include .env #https://stackoverflow.com/a/8346208
run: configure source_env
	dep ensure
	go run main.go

# operational stuffs
configure:
	@echo "Running configure"
	@test ! -s .env && cp -iv .env.sample .env || echo '.env file already there'

source_env:
	@test -s .env && export $(shell sed 's/=.*//' .env) &> /dev/null

list_jobs:
	@go run main.go list jobs
