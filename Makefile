
install-swags:
	go get -u github.com/swaggo/swag/cmd/swag
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files

swag-init:
	go run github.com/swaggo/swag/cmd/swag init

run:
	go run main.go

