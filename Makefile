include base.mk

routes_mkfile_dir := $(CURRDIR)

check-dependencies:
	@$(call check-dependency,go)
	@$(call check-dependency,jq)
	@$(call check-dependency,deck)
	@$(call check-dependency,docker)

test: check-dependencies
	@go test -v ./...

build: check-dependencies
	@go generate ./...
	@go build .

build-docker:
	@docker build -t kongair/routes:latest .

run: check-dependencies build
	@./routes ${KONG_AIR_ROUTES_PORT}

docker: build-docker
	@docker run -d --name kongair-routes -p ${KONG_AIR_ROUTES_PORT}:${KONG_AIR_ROUTES_PORT} kongair/routes:latest

kill-docker:
	-@docker stop kong-air-routes-svc
	-@docker rm kong-air-routes-svc
	@if [ $$? -ne 0 ]; then $(call echo_fail,Failed to kill the docker containers); exit 1; else $(call echo_pass,Killed the docker container); fi

