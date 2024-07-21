require("@nomiclabs/hardhat-ethers");
require("@nomiclabs/hardhat-waffle");
module.exports = {
  solidity: {
    version: "0.8.0",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    hardhat: {
      chainId: 1337,
      gas: 12000000, // Increase the gas limit here
      blockGasLimit: 12000000, // Optional: set block gas limit
    },
  },
};

