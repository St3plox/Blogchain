GOLANG 			:= golang:1.22
ALPINE 			:= alpine:3.19
EXPOSE_PORT 	:= 3000
INTERNAL_PORT 	:= 3000
BACKEND_NAME 	:= blockchain-back-api
BASE_IMAGE_NAME := localhost/tveu/blogchain
VERSION 		:= 0.0.1
BACKEND_IMAGE 	:= $(BASE_IMAGE_NAME)/$(BACKEND_NAME):$(VERSION)

npm-tools:
	npm install --save-dev @nomiclabs/hardhat-ethers ethers waffle chai

dev-docker:
	docker pull $(ALPINE)
	docker pull $(GOLANG)

gen-private:
	cd zarf
	mkdir keys
	cd keys
	openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:4096

run-local:
	go run app/backend/user-service/main.go

build-image:
	docker build -t $(BACKEND_IMAGE) -f zarf/docker/Dockerfile .

run:
	docker run -d -p $(EXPOSE_PORT):$(INTERNAL_PORT) --name $(BACKEND_NAME) $(BACKEND_IMAGE)

logs:
	docker logs $(BACKEND_NAME)

stop:
	docker stop $(BACKEND_NAME)
	docker rm $(BACKEND_NAME)

run-tests:
	go test ./foundation/blockchain 

launch-mongo:
	docker run -d -p 27017:27017 --name mongodb mongodb/mongodb-community-server:6.0-ubi8

stop-mongo:
	docker stop mongodb
	docker rm mongodb