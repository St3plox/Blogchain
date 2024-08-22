GOLANG 			:= golang:1.22
ALPINE 			:= alpine:3.19
EXPOSE_PORT 	:= 3000
INTERNAL_PORT 	:= 3000
BACKEND_NAME 	:= blockchain-back-api
BASE_IMAGE_NAME := localhost/tveu/blogchain
VERSION 		:= 0.0.1
BACKEND_IMAGE 	:= $(BASE_IMAGE_NAME)/$(BACKEND_NAME):$(VERSION)
REDIS_NAME		:= redis
REDIS_PORT		:= 6379


dev-tools:
	 go install github.com/swaggo/swag/cmd/swag@latest
	 go install github.com/ethereum/go-ethereum/cmd/abigen@latest

npm-tools:
	npm install --save-dev @nomiclabs/hardhat-ethers ethers waffle chai

dev-docker:
	docker pull $(ALPINE)
	docker pull $(GOLANG)

gen-private:
	mkdir zarf/keys
	openssl genpkey -algorithm RSA -out zarf/keys/private_key.pem -pkeyopt rsa_keygen_bits:4096

gen-docs:
	swag init -g app/backend/user-service/main.go --parseDependency --parseInternal

run-local:
	go run app/backend/user-service/main.go

build-image-service:
	docker build -t $(BACKEND_IMAGE) -f zarf/docker/Dockerfile .

run-service:
	docker run -d -p $(EXPOSE_PORT):$(INTERNAL_PORT) --name $(BACKEND_NAME) $(BACKEND_IMAGE)

run-hardhat:
	docker build -t hardhat-node -f zarf/docker/hardhat/Dockerfile .
	docker run -d -p 8545:8545 --name hardhat-node hardhat-node

stop-hardhat:
	docker stop hardhat-node
	docker rm hardhat-node

logs-service:
	docker logs $(BACKEND_NAME)

stop-service:
	docker stop $(BACKEND_NAME)
	docker rm $(BACKEND_NAME)
 

launch-hardhat:
	npx hardhat node

launch-mongo:
	docker run -d -p 27017:27017 --name mongodb mongodb/mongodb-community-server:6.0-ubi8

stop-mongo:
	docker stop mongodb
	docker rm mongodb

launch-redis:
	docker run -d --name $(REDIS_NAME) -p $(REDIS_PORT):$(REDIS_PORT)  redis/redis-stack-server:latest

stop-redis:
	docker stop $(REDIS_NAME)
	docker rm &(REDIS_NAME)

run-front:
	cd app/frontend/blogchain-vue
	npm run dev

front-deps:
	cd app/frontend/blogchain-vue
	npm install

solc-compile:
	 solc --overwrite --abi --bin -o contracts/bin contracts/PostStorage.sol

generate-contract:
	abigen --abi=contracts/bin/PostStorage.abi --bin=contracts/bin/PostStorage.bin --pkg=contract --out=foundation/blockchain/contract/post_storage.go



#------------------------------------------------------------------------------------------
#TESTING

test-blockchain-auth:
	go test ./foundation/blockchain/auth

test-blockchain-contract:
	go test ./foundation/blockchain 



#launch hardhat node before that 
test-foundation:
	make test-blockchain-auth
	make test-blockchain-contract
	

test-cachestore:
	go test ./foundation/cachestore

test-userdb:
	 go test  ./business/core/user/userdb

test-user:
	go test ./business/core/user
