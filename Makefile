.PHONY: tailwind-watch
tailwind-watch:
	npx @tailwindcss/cli -i ./assets/css/main.css -o ./public/assets/css/main.css --watch