include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d shop-postgres shop-redis

env-down:
	@docker compose down shop-postgres shop-redis

env-cleanup:
	@read -p "Очисти всі volume файли? [y/N]" ans; \
	if [ "$$ans" == "y" ]; then \
	  docker compose down shop-postgres shop-redis && \
	  sudo rm -rf out/pgdata && \
	  sudo rm -rf out/redis_data && \
	  echo "Volume файли очищені!"; \
	else \
	  echo "Очистка середовища скасована"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
  		echo "Відсутній seq"; \
  		exit 1; \
  	fi;
	docker compose run --rm shop-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
  		echo "Відстуній action"; \
  		exit 1; \
  	fi;
	docker compose run --rm shop-migrate \
    	-path /migrations \
    	-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@shop-postgres:5432/${POSTGRES_DB}?sslmode=disable \
    	"$(action)"

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder