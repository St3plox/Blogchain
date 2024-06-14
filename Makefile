GOLANG := golang:1.22

npm-tools:
	npm install --save-dev @nomiclabs/hardhat-ethers ethers waffle chai

gen-private:
	cd zarf
	mkdir keys
	cd keys
	openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:4096

run-local:
	go run app/backend/user-service/main.go

run-tests:
	go test ./foundation/blockchain/account_test.go

launch-mongo:
	docker run -d -p 27017:27017 --name mongodb mongodb/mongodb-community-server:6.0-ubi8