.PHONY: client server

test-client:
	cd ./client && clear && npm run test

client:
	cd ./client && npm run dev

server:
	cd ./server && nodemon --exec "go run" main.go