.PHONY: tailwind-watch
tailwind-watch:
	npx @tailwindcss/cli -i ./assets/css/main.css -o ./public/assets/css/main.css --watch

.PHONY: tsc-watch
tsc-watch:
	npx tsc --watch

.PHONY: compose-dev-build
compose-dev-build:
	docker compose -f compose.dev.yaml build

.PHONY: compose-dev-up
compose-dev-up:
	docker compose -f compose.dev.yaml up -d

.PHONY: compose-dev-down
compose-dev-down:
	docker compose -f compose.dev.yaml down -d

.PHONY: shell-app
shell-app:
	docker exec -it lesfraternels_website_app sh
