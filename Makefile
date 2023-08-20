build: 
	@echo -e "\033[1;34m[*] Building\033[0m\n"
	@go build -o ./bin/app
	@cd ./frontend/svelte-app &&  npm install && npm run build
	@echo -e "\033[1;32m[+] Done Building\033[0m\n"

run: build
	@echo -e "\033[1;34m[*] Running Server\033[0m\n"
	@./bin/app
