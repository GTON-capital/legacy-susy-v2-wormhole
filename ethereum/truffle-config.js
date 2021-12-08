require('dotenv').config({ path: ".env" });
const HDWalletProvider = require('@truffle/hdwallet-provider');

module.exports = {
    networks: {
        development: {
            host: "127.0.0.1",
            port: 7545,
            network_id: "*",
        },
        mainnet: {
            provider: () => new HDWalletProvider(process.env.MNEMONIC, `https://mainnet.infura.io/v3/`+process.env.INFURA_KEY),
            network_id: 1,
            gas: 10000000,
            gasPrice: 101000000000,
            confirmations: 1,
            timeoutBlocks: 200,
            skipDryRun: false
        },
        rinkeby: {
            provider: () => new HDWalletProvider(process.env.MNEMONIC, `https://rinkeby.infura.io/v3/`+process.env.INFURA_KEY),
            network_id: 4,
            gas: 5500000,
            confirmations: 0,
            timeoutBlocks: 300,
            skipDryRun: true
        },
        ropsten: {
            provider: () => new HDWalletProvider(process.env.MNEMONIC, `https://ropsten.infura.io/v3/`+process.env.INFURA_KEY),
            network_id: 3,
            gas: 5500000,
            confirmations: 0,
            timeoutBlocks: 300,
            skipDryRun: true
        },
        goerli: {
            provider: () => {
                return new HDWalletProvider(process.env.MNEMONIC, 'https://goerli.infura.io/v3/'+process.env.INFURA_KEY)
            },
            network_id: '5',
            gas: 4465030,
            gasPrice: 10000000000,
        },
        fantom: {
            provider: () => {
                // return new HDWalletProvider(process.env.MNEMONIC, 'https://apis.ankr.com/6052791850e6426392593b0ddba45bf5/d37735e535d9d051230799cae45aeb6a/polygon/full/main')
                // return new HDWalletProvider(process.env.MNEMONIC, 'https://polygon-rpc.com')
                return new HDWalletProvider(process.env.MNEMONIC, 'https://apis.ankr.com/c3f382db4c55497b81bf3feb1e9a8499/d37735e535d9d051230799cae45aeb6a/fantom/full/main')
            },
            // network_id: 137,
            // gasPrice: 70 * 1e9,
            // confirmations: 1,
            // skipDryRun: false,
            // timeoutBlocks: 200,
            network_id: 250,
            // gas: 70000000,
            // gasPrice: 8000000000,
            gasPrice: 100 * 1e9,
            confirmations: 1,
            skipDryRun: false,
            timeoutBlocks: 200,
        },
        polygon: {
            provider: () => {
                // return new HDWalletProvider(process.env.MNEMONIC, 'https://apis.ankr.com/6052791850e6426392593b0ddba45bf5/d37735e535d9d051230799cae45aeb6a/polygon/full/main')
                // return new HDWalletProvider(process.env.MNEMONIC, 'https://polygon-rpc.com')
                return new HDWalletProvider(process.env.MNEMONIC, 'https://apis.ankr.com/6052791850e6426392593b0ddba45bf5/d37735e535d9d051230799cae45aeb6a/polygon/full/main')
            },
            // network_id: 137,
            // gasPrice: 70 * 1e9,
            // confirmations: 1,
            // skipDryRun: false,
            // timeoutBlocks: 200,
            network_id: 137,
            // gas: 70000000,
            // gasPrice: 8000000000,
            // gasPrice: 25 * 1e9,
            confirmations: 4,
            skipDryRun: false,
            timeoutBlocks: 200,
        },
        binance: {
            provider: () => {
                return new HDWalletProvider(process.env.MNEMONIC, 'https://arb:arb@apis.ankr.com/6017ce5a3b86463c8b15279cc25d74b3/d37735e535d9d051230799cae45aeb6a/binance/full/main')
            },
            network_id: 56,
            // gas: 70000000,
            // gasPrice: 8000000000,
            gasPrice: 5 * 1e9,
            confirmations: 1,
            skipDryRun: false,
            timeoutBlocks: 200,
        },
    },

    compilers: {
        solc: {
            version: "0.8.4",
            settings: {
                 optimizer: {
                   enabled: true,
                   runs: 200
                 },
            }
        }
    },

    plugins: [
        "@chainsafe/truffle-plugin-abigen",
        "truffle-plugin-verify"
    ],

    api_keys: {
        etherscan: 'K5KZ6G1N1IEGDJNDQG9VAGW3SBW2EZYN4N'
    },
};
