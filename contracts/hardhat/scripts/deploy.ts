import { ethers } from "hardhat";
import * as process from "process";

async function main() {
  const [deployer] = await ethers.getSigners();
  console.log("Deploying contracts with the account:", deployer.address);
  const token = await ethers.getContractFactory("GToken");
    // const gToken = await token.deploy("GToken", "GTP", "1000000000"); // eth contract
    const gToken = await token.deploy("GToken", "GTP", "1000000000"); // bsc contract  -> use 0.8.19 because PUSH0 opcode only available for Eth net
  await gToken.waitForDeployment();
  const address = await gToken.getAddress();
  console.log(`GToken deployed to ${gToken.target} at address ${address}`);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
});
