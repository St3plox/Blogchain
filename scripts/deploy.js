const fs = require('fs');
const path = require('path');
const { ethers } = require('hardhat');

async function main() {
  // Get all contract files from the contracts directory
  const contractsDir = path.join(__dirname, '../contracts');
  const contractFiles = fs.readdirSync(contractsDir).filter(file => file.endsWith('.sol'));

  // Create an object to store the contract ABIs and addresses
  const contractData = {};

  // Loop through each contract file and deploy it
  for (const file of contractFiles) {
    const contractName = file.replace('.sol', '');
    console.log(`Deploying ${contractName}...`);

    // Compile and deploy the contract
    const ContractFactory = await ethers.getContractFactory(contractName);
    const contract = await ContractFactory.deploy();
    await contract.deployed();

    console.log(`${contractName} deployed at ${contract.address}`);

    // Get the contract ABI
    const artifact = await hre.artifacts.readArtifact(contractName);
    const { abi } = artifact;

    // Save the ABI and address to the contractData object
    contractData[contractName] = {
      address: contract.address,
      abi: abi
    }; 
  }

  // Write the contract data to a JSON file in the /contract/cfg directory
  const cfgDir = path.join(__dirname, '../contracts/cfg');
  if (!fs.existsSync(cfgDir)) {
    fs.mkdirSync(cfgDir, { recursive: true });
  }

  const outputFile = path.join(cfgDir, 'deployedContracts.json');
  fs.writeFileSync(outputFile, JSON.stringify(contractData, null, 2));

  console.log(`Contract ABIs and addresses saved to ${outputFile}`);
}

// Execute the main function
main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error);
    process.exit(1);
  });
