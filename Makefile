.PHONY: client server

client:
	cd ./client && npm run dev

server:
	cd ./server && nodemon --exec "go run" main.go