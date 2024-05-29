.PHONY: serve build tailwindcss air

build:
	@./dev/tailwindcss/bin -i tailwind.css -o web/public/css/main.css --minify
	@./dev/templ/bin generate
	@go build -o ./tmp/main .

dev: air open-browser tailwindcss
	
tailwindcss:
	@./dev/tailwindcss/bin -i tailwind.css -o ./web/public/css/main.css --minify --watch

air:
	@./dev/air/bin &

open-browser:
	@(sleep 2; xdg-open "http://localhost:3333") &
