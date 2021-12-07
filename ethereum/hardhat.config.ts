import "@typechain/hardhat";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";
import "hardhat-abi-exporter";

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
export default {
  networks: {
    localhost: {
      url: "http://127.0.0.1:7545",
    },
    avax: {
      url: "https://api.avax.network/ext/bc/C/rpc",
      accounts: [],
    },
    harmony: {
      url: "https://api.harmony.one",
      accounts: [],
    },
    fantom: {
      url: "https://rpcapi-tracing.fantom.network",
      accounts: [],
    },
    polygon: {
      url: "https://apis.ankr.com/6052791850e6426392593b0ddba45bf5/d37735e535d9d051230799cae45aeb6a/polygon/full/main",
      accounts: [],
    },
  },
  solidity: {
    compilers: [
      {
        version: "0.8.4",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
      {
        version: "0.6.6",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
      {
        version: "0.6.12",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
      {
        version: "0.5.16",
        settings: {
          optimizer: {
            enabled: true,
            runs: 200,
          },
        },
      },
    ],
  },
  abiExporter: {
    clear: true,
    flat: true,
    spacing: 2,
  },
  mocha: {
    timeout: "100000000000000",
  },
};
