up:
	@docker compose -f deploy/compose.yaml up -d

down:
	@docker compose -f deploy/compose.yaml down

.PHONY: up down