require("@nomiclabs/hardhat-ethers");
require("@nomiclabs/hardhat-waffle");
require("hardhat-abigen");
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
    },
  },
  abigen: {
    outDir: "abi",            // The output directory for generated ABI files (default: "abi")
    inDir: "contracts",       // The input directory containing your contract files (default: "contracts")
    includeContracts: ["*"],  // An array of contract patterns to include in the generate ABIs (default: ["*"])
    excludeContracts: [],     // An array of contract patterns to exclude from the generate ABIs (default: [])
    space: 2,                 // The number of spaces to use for indentation in the generated ABIs (default: 2)
    autoCompile: true         // Whether to automatically compile contracts before generating ABIs (default: true)
  },
};
