.PHONY: help

## help:	List all commands with descriptions
help: Makefile
	@sed -n "s/^##//p" $<

## run-app:  		Run app using go run command | Usage: APP_SECRET=<app_secret> make run-app
run-app: key_gen
	@export $(APP_SECRET)
	@go run main.go --host=127.0.0.1 --port=5000 --private-key-path=private.pem --public-key-path=public.pem

## key_gen:		Generates RSA public, private key pair
key_gen: gen_rsa

gen_rsa:
	@openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
	@openssl rsa -in private.pem -pubout -out public.pem

## test: 			Run tests
test:
	go test ./...