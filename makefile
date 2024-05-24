.PHONY: serve build tailwindcss air

build:
	@./dev/tailwindcss/bin -o web/public/css/main.css --minify
	@./dev/templ/bin generate
	@go build -o ./tmp/main .

dev: air tailwindcss
	
tailwindcss:
	@./dev/tailwindcss/bin -o ./web/public/css/main.css --minify --watch

air:
	@./dev/air/bin &
