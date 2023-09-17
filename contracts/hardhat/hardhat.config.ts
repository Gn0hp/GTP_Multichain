// look through this sample at https://github.com/ethereum-optimism/contracts/blob/master/hardhat.config.ts
import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
require('dotenv').config();

const PRIVATE_KEY = process.env.PRIVATE_KEY || "";
const config: HardhatUserConfig = {
  defaultNetwork: "sepolia",
  solidity: {
    version: "0.8.19",
    settings: {
      optimizer: {
        enabled: false,
        runs: 200,
      }
    }
  },
  networks: {
    sepolia: {
      url: "https://sepolia.infura.io/v3/32d2aa75e85d41ebbf89dcd05b9d63a1",
      chainId: 11155111,
      accounts: [PRIVATE_KEY]
    },
    bsc_testnet: {
      url: "https://data-seed-prebsc-1-s3.binance.org:8545/",
      chainId: 97,
      gasPrice: 8800000000,
      accounts: [PRIVATE_KEY]
    }
  },
};

export default config;
