GOLANG := golang:1.22

npm-tools:
	npm install --save-dev @nomiclabs/hardhat-ethers ethers waffle chai

run-local:
	go run app/backend/user-service/main.go