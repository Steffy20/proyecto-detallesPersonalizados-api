.PHONY: up stop down logs

up:
	docker compose up --build -d

stop:
	docker compose down

down:
	docker compose down -v

logs:
	docker compose logs -f api
