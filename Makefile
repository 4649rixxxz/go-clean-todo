build:
	docker compose build --no-cache --force-rm
up:
	docker compose up -d 
create-project:
	@make build
	@make up
destroy:
	docker compose down --rmi all --volumes --remove-orphans
start:
	docker compose start
stop:
	docker compose stop
.PHONY: api
api:
	docker compose exec api /bin/sh
db:
	docker compose exec db bash